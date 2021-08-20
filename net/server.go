package net

import (
	"fmt"
	"net"
	"zinx/iFace"
)

//Server IServer 接口的实现，定义了一个 Server 的服务器模块
type Server struct {
	Name string //服务器的名称
	IPVersion string //服务器绑定的 IP 版本
	IP string //服务器监听的 IP
	Port uint32 //服务器监听的端口
}

func NewServer(name string) iFace.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8080,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP:%s Port:%d, is starting\n", s.IP, s.Port)

	go func() {
		//获取待创建的服务器的 TCP 地址
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp address default,error:", err)
		}

		//监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listen error:", err)
		}
		// 监听成功
		fmt.Println("start Zinx server ", s.Name, " success, now listening...")

		//启动 server 网络连接业务
		for {
			//阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp error:", err)
				continue
			}

			// 最大回显 512 字节的业务
			go func() {
				// 不停从客户端循环的获取数据
				for {
					buf := make([]byte, 512)
					//这里最多只能读取缓冲区 len(buf) （不是 cap(buf)）字节大小的客户端内容
					readLen, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read client content error:", err)
						continue
					}

					// 回显
					if _, err := conn.Write(buf[:readLen]); err != nil {
						fmt.Println("write back buf error:", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Server() {
	s.Start()

	//阻塞
	select{}
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name " , s.Name)
}
