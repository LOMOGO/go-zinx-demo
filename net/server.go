package net

import (
	"fmt"
	"net"
	"zinx/iFace"
	"zinx/utils"
)

//Server IServer 接口的实现，定义了一个 Server 的服务器模块
type Server struct {
	Name string //服务器的名称
	IPVersion string //服务器绑定的 IP 版本
	IP string //服务器监听的 IP
	Port uint32 //服务器监听的端口
	Router iFace.IRouter //当前 server 由用户绑定的回调 router
}

func NewServer(config utils.Config) iFace.IServer {
	if config.IP == "" {
		panic("没有给出服务器 IP")
	}
	utils.GlobalConfig.IP = config.IP

	if config.Port == 0 {
		panic("没有给出服务器端口号")
	}
	utils.GlobalConfig.Port = config.Port

	if config.Name != utils.GlobalConfig.Name && config.Name != "" {
		utils.GlobalConfig.Name = config.Name
	}
	if config.MaxConn != 0 && config.MaxConn != utils.GlobalConfig.MaxConn {
		utils.GlobalConfig.MaxConn = config.MaxConn
	}
	if config.MaxPackage != 0 && config.MaxPackage != utils.GlobalConfig.MaxPackage {
		utils.GlobalConfig.MaxPackage = config.MaxPackage
	}

	s := &Server{
		Name:      utils.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalConfig.IP,
		Port:      utils.GlobalConfig.Port,
		Router: nil,
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

		var cid uint32 = 0

		//启动 server 网络连接业务
		for {
			//阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp error:", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Server() {
	s.Start()

	//阻塞以防止该 goroutine 退出
	select{}
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name " , s.Name)
}

func (s *Server) AddRouter(router iFace.IRouter) {
	s.Router = router
	fmt.Println("add router success")
}
