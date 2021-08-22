package iFace

//IServer 定义一个服务器接口
type IServer interface {
	//Start 启动服务器
	Start()
	//Server 运行服务器
	Server()
	//Stop 停止服务器
	Stop()
	//AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
	AddRouter(router IRouter)
}
