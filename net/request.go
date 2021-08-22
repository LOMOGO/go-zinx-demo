package net

import "zinx/iFace"

type Request struct {
	conn iFace.IConnection
	data []byte
}

func (r *Request) GetConnection() iFace.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
