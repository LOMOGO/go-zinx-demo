package iFace

//IRequest 将客户端请求的连接信息和发送过来的请求数据都包裹在 Request 中
type IRequest interface {
	GetConnection() IConnection //获取请求的连接信息
	GetData() []byte //获取请求发送来的数据
}
