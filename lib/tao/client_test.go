package tao

import (
	"context"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"log"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	tn := &TempNotify{}
	tao.Register(tn.MessageNumber(), nil, func(ctx context.Context, closer tao.WriteCloser) {
		notify := tao.MessageFromContext(ctx).(*TempNotify)
		log.Println("==>   ", len(notify.Zones))
	})
	c, err := net.Dial("tcp", "192.168.0.86:17083")
	if err != nil {
		holmes.Fatalln(err)
	}
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		log.Println("on connect")
		err := conn.Write(&Connected{})
		if err != nil {
			log.Fatal("Connected ", err)
		}
		//err = conn.Write(&DeviceRequest{})
		//if err != nil {
		//	log.Fatal("DeviceRequest ",err)
		//}
		return true
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		log.Println("on error")
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		log.Println("on close")
	})

	onMessage := tao.OnMessageOption(func(msg tao.Message, conn tao.WriteCloser) {
		msg.MessageNumber()
	})

	conn := tao.NewClientConn(0, c, onConnect, onError, onClose, onMessage, tao.ReconnectOption())
	conn.Start()
	select {}
}
