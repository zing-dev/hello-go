package main

import "encoding/json"

type Response struct {
	Success bool
	Err     string
}

//打开继电器请求参数
type RequestOpen struct {
	Relay []bool
}

//打开继电器请求参数
type RequestClose struct {
	Relay []bool
}

//打开继电器请求参数
type RequestReset struct{}

type MsgId byte

const (
	ConnectId    MsgId = 0   //连接
	DisconnectId MsgId = 1   //关闭连接
	OpenId       MsgId = 2   //打开继电器
	CloseId      MsgId = 3   //关闭继电器
	ResetId      MsgId = 4   //重置继电器
	HeartBeatId  MsgId = 250 //心跳
	IllegalId    MsgId = 255 //非法请求
)

func (m MsgId) ToByte() byte {
	return byte(m)
}

func (m MsgId) ToBytes() []byte {
	switch m {
	case OpenId:
		d, _ := json.Marshal(RequestOpen{Relay: T})
		return d
	case CloseId:
		d, _ := json.Marshal(RequestClose{Relay: T})
		return d
	case ResetId:
		d, _ := json.Marshal(RequestReset{})
		return d
	case IllegalId:
		fallthrough
	default:
		return nil
	}
}
