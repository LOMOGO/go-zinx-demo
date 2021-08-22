package net

import (
	"fmt"
	"net"
	"testing"
	"time"
	"zinx/iFace"
)

type PingRouter struct {
	BaseRouter //匿名字段
}

func (p *PingRouter) Handle(request iFace.IRequest) {
	fmt.Println("call router handle func")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("[server] ping..."))
	if err != nil {
		fmt.Println("[err] server ping error:", err)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("lomogo")
	s.AddRouter(&PingRouter{})
	go client()
	s.Server()
}

func client() {
	fmt.Println("client test start...")
	time.Sleep(3*time.Second)

	//客户端与服务器建立连接
	conn, err :=  net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start error:", err)
		return
	}

	for {
		//客户端间隔一秒发送数据
		_, err := conn.Write([]byte("yq @ br"))
		if err != nil {
			fmt.Println("client write content error:", err)
			continue
		}

		//客户端读取服务端发送回来的数据
		buf := make([]byte, 512)
		cntLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read content error:", err)
			continue
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf,  cntLen)

		time.Sleep(1*time.Second)
	}
}
