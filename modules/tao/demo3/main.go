package main

import (
	"github.com/leesper/holmes"
	"log"
	"net"
	"test-tao/msg"
	"test-tao/tao"
	"time"
)

func main() {
	defer holmes.Start().Stop()
	temp := msg.TempRequest{
		Request: &msg.ZoneTempNotify{},
	}
	sign := msg.SignResponse{
		Request: &msg.TempSignalNotify{},
	}
	tao.Register(temp.MessageNumber(), temp.Unmarshaler, temp.Handle)
	tao.Register(sign.MessageNumber(), sign.Unmarshaler, sign.Handle)
	c, err := net.Dial("tcp", "192.168.0.215:17083")
	if err != nil {
		holmes.Fatalln(err)
	}
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		conn.Write(&msg.DeviceRequest{Request: &msg.SetDeviceRequest{
			ZoneTempNotifyEnable:    true,
			ZoneAlarmNotifyEnable:   false,
			FiberStatusNotifyEnable: false,
			TempSignalNotifyEnable:  false,
		}})
		go func(conn tao.WriteCloser) {
			for {
				conn.Write(&msg.PingRequest{})
				time.Sleep(time.Second * 10)
			}
		}(conn)
		return true
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		log.Println("on error")
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		log.Println("on close")
	})

	onMessage := tao.OnMessageOption(func(msg tao.Message, conn tao.WriteCloser) {
		log.Println(msg.MessageNumber())
	})

	conn := tao.NewClientConn(0, c, onConnect, onError, onClose, onMessage, tao.ReconnectOption())
	conn.Start()
	select {}

}
