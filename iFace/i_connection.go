package iFace

import "net"

type IConnection interface {
	//Start 启动连接，让当前连接开始工作
	Start()
	//Stop 结束当前连接状态
	Stop()
	//GetTCPConnection 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//GetConnID 获取当前的连接ID
	GetConnID() uint32
	//RemoteAddr 获取远程客户端地址信息
	RemoteAddr() net.Addr
}
