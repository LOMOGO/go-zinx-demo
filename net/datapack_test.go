package net

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	go client1()
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	fmt.Println("[server] start...")

	if err != nil {
		fmt.Println("server listen error:", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("accept conn error:", err)
		}

		go func(conn net.Conn) {
			dp := NewDataPack()

			for {
				headData := make([]byte, dp.GetMsgHeadLen())
				if _, err := io.ReadFull(conn, headData); err != nil {
					fmt.Println("read head error:", err)
				}

				msgHead, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("server unpack head error:", err)
				}

				if msgHead.GetMsgDataLen() > 0 {
					msg := msgHead.(*Message)
					msg.data = make([]byte, msg.GetMsgDataLen())

					if _, err := io.ReadFull(conn, msg.data); err != nil {
						fmt.Println("server unpack data error:", err)
						return
					}

					fmt.Println("receive msg ID:", msg.GetMsgID(), ", len:", msg.GetMsgDataLen(), ", data:", string(msg.GetMsgData()))
				}
			}
		}(conn)
	}
}

func client1() {
	fmt.Println("[client] start...")
	time.Sleep(3*time.Second)
	//客户端goroutine，负责模拟粘包的数据，然后进行发送
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	//创建一个封包对象 dp
	dp := NewDataPack()

	//封装一个msg1包
	msg1 := &Message{
		id:      0,
		len: 5,
		data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
		return
	}

	msg2 := &Message{
		id:      1,
		len: 7,
		data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client temp msg2 err:", err)
		return
	}

	//将sendData1，和 sendData2 拼接一起，组成粘包
	sendData1 = append(sendData1, sendData2...)

	//向服务器端写数据
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}
