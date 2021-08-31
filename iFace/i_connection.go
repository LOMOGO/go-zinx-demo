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
	//SendMsg 直接将 message 数据发送给远程 TCP 客户端
	SendMsg(msgID uint32, data []byte) error
}
