package net

import (
	"fmt"
	"net"
	"zinx/iFace"
)

type Connection struct {
	Conn *net.TCPConn //当前 TCP 连接的 socket 套接字
	ConnId uint32 //当前连接的 ID （Session ID）， ID 是全局唯一的
	isClosed bool //当前连接的关闭状态
	handleAPI iFace.HandFunc //该连接的处理方法 api
	ExitBuffChan chan bool //告知该连接已经停止/退出的 channel（信号通知）
}

//NewConnection 新建一个连接
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi iFace.HandFunc) iFace.IConnection {
	c := &Connection{
		Conn: conn,
		ConnId: connID,
		isClosed: false,
		handleAPI: callbackApi,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

//StartReader 处理 conn 接收到的客户端数据的 goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit")
	defer close(c.ExitBuffChan)

	for {
		//最大读取 512 字节到 buf 中
		buf := make([]byte, 512)
		cntLen, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error:", err)
			c.ExitBuffChan <- true
			continue
		}

		//调用当前连接业务（执行的是与当前连接绑定的 handle 方法）
		if err := c.handleAPI(c.Conn, buf, cntLen); err != nil {
			fmt.Println("connID:", c.ConnId, " handle is error")
			c.ExitBuffChan <- true
			return
		}
	}
}

//Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {

	//开启处理该连接读取到客户端数据之后的请求业务
	go c.StartReader()

	//阻塞，直到得到通知退出
	for {
		select {
		case <- c.ExitBuffChan:
			return
		}
	}
}

//Stop 停止连接，结束当前的连接状态
func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭 socket 连接
	c.Conn.Close()

	//通知读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true

	//关闭该连接的全部通道
	close(c.ExitBuffChan)
}

//GetTCPConnection 获取连接的 TCP socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//GetConnID 获取会话 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

//RemoteAddr 获取客户端的地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
