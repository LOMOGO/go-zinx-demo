package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	s := NewServer("lomogo")
	go client()
	s.Server()
}

func client() {
	fmt.Println("client test start...")
	time.Sleep(3*time.Second)

	conn, err :=  net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start error:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("yq @ br"))
		if err != nil {
			fmt.Println("client write content error:", err)
			continue
		}

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
