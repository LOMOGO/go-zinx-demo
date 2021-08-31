package iFace

//IMessage 将请求的一个消息封装到 message 中，message 实现了 IMessage
type IMessage interface {
	GetMsgDataLen() uint32 //获取一个 message 中 data 的长度
	GetMsgID() uint32      //获取一个 message 的 ID
	GetMsgData() []byte    //获取一个 message 的内容数据

	SetMsgLen(length uint32) //设置一个 message 的长度
	SetMsgID(id uint32) //设置一个 message 的 ID
	SetMsgData(data []byte) //设置一个 message 的内容
}
