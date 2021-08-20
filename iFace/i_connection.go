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

//HandFunc 如果想要指定一个 conn 的处理业务，只要定义一个 HandFunc 类型的函数，然后与该 conn 绑定即可。
//其中第一个参数是 socket 原生连接， 第二个参数是客户端请求的数据，第三个参数是客户端请求的数据长度
type HandFunc func(*net.TCPConn, []byte, int) error
