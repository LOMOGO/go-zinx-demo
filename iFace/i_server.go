package iFace

//IServer 定义一个服务器接口
type IServer interface {
	//Start 启动服务器
	Start()
	//Server 运行服务器
	Server()
	//Stop 停止服务器
	Stop()
}
