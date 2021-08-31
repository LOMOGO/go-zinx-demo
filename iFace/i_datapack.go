package iFace

type IDataPack interface {
	GetMsgHeadLen() uint32 //获取一条完整的消息的头部所占字节大小
	Pack(msg IMessage) ([]byte, error) //封包操作，将一条消息封装
	UnPack([]byte) (IMessage, error) //拆包操作，将流中的一条消息解析出来
}
