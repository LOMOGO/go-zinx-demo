package net

import "zinx/iFace"

//BaseRouter 实现 router 时，先嵌入这个基类，然后再根据需要对基类中的方法进行重写
type BaseRouter struct {}

//这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化

func (bs *BaseRouter) PreHandle(request iFace.IRequest) {}
func (bs *BaseRouter) Handle(request iFace.IRequest) {}
func (bs *BaseRouter) PostHandle(request iFace.IRequest) {}
