package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/iFace"
	"zinx/utils"
)

type DataPack struct {}

const msgHeadLen = 8 //一条完整的消息中头部所占的字节大小

func (d *DataPack) GetMsgHeadLen() uint32 {
	return msgHeadLen
}

func NewDataPack() iFace.IDataPack {
	return &DataPack{}
}

//Pack 将一条消息进行封包操作
func (d *DataPack) Pack(msg iFace.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

//UnPack 将一条消息进行拆包操作
func (d *DataPack) UnPack(binaryData []byte) (iFace.IMessage, error) {
	dataBuf := bytes.NewReader(binaryData)

	var msg Message

	//由于包头部分的长度是固定的，包体（data）部分的长度是不固定的，因此首先需要读取出固定长度的字节：len 和 id
	//在读取了以上固定长度的部分后，需要根据读出的 len 来读取 data,这一步骤交由下一个函数执行
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.len); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}

	//如果数据的长度超出最大包限制，将报错
	if msg.GetMsgDataLen() > utils.GlobalConfig.MaxPackage {
		return nil, errors.New("receive msg Data too large")
	}

	return &msg, nil
}
