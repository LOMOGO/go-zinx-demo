package iFace

//IRouter 这里的路由是框架使用者将连接自定的处理业务方法和连接绑定在一起
//其中的 IRequest 则包含用该链接的链接信息和该链接的请求数据信息（路由可以理解为将不同的连接和处理该连接的业务方法绑定在一起）
type IRouter interface {
	PreHandle(request IRequest) // 在处理 conn 业务之前的钩子方法（所谓钩子方法也就是基方法，可以等待别的方法继承完善基方法）
	Handle(request IRequest) // 处理 conn 业务的钩子方法
	PostHandle(request IRequest) // 在处理 conn 业务之后的钩子方法
}
