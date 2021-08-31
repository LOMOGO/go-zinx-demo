package net

import "zinx/iFace"

type Request struct {
	conn iFace.IConnection
	data iFace.IMessage
}

func (r *Request) GetConnection() iFace.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data.GetMsgData()
}

func (r *Request) GetID() uint32 {
	return r.data.GetMsgID()
}
