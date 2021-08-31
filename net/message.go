package net

//Message Message 结构体实现了 IMessage 接口
type Message struct {
	len  uint32 //一个 message 的长度
	id   uint32 //一个 message 的ID
	data []byte //一个 message 的内容
}

func (m *Message) GetMsgDataLen() uint32 {
	return m.len
}

func (m *Message) GetMsgID() uint32 {
	return m.id
}

func (m *Message) GetMsgData() []byte {
	return m.data
}

func (m *Message) SetMsgLen(length uint32) {
	m.len = length
}

func (m *Message) SetMsgID(id uint32){
	m.id = id
}

func (m *Message) SetMsgData(data []byte) {
	m.data = data
}
