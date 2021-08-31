package net

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"zinx/iFace"
	"zinx/utils"
)

type PingRouter struct {
	BaseRouter //匿名字段
}

func (p *PingRouter) Handle(request iFace.IRequest) {
	fmt.Println("call router handle func")
	//先读取客户端发送来的数据，再回写 pong
	fmt.Println("[server] receive from client msg ID:", request.GetID(), "Data:", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("[server] pong...pong...pong"))
	if err != nil {
		fmt.Println("[err] server ping error:", err)
	}
}

func TestServer(t *testing.T) {
	s := NewServer(utils.DefaultConfig)
	s.AddRouter(&PingRouter{})
	go client()
	s.Server()
}

//client 模拟客户端
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
		//发封包message消息
		dp := NewDataPack()
		data := []byte("[client] ping...ping...ping")
		msg, _ := dp.Pack(&Message{
			len:  uint32(len(data)),
			id:   0,
			data: data,
		})
		_, err := conn.Write(msg)
		if err !=nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dp.GetMsgHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetMsgDataLen() > 0 {
			data := make([]byte, msgHead.GetMsgDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			msgHead.SetMsgData(data)

			fmt.Println("[client] receive form server msg: ID=", msgHead.GetMsgID(), ", len=", msgHead.GetMsgDataLen(), ", data=", string(msgHead.GetMsgData()))
		}

		time.Sleep(1*time.Second)
	}
}
