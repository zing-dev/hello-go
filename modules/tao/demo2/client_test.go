package demo2

import (
	"fmt"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"net"
	"testing"
	"time"
)

func register() {
	u := new(User)
	tao.Register(u.MessageNumber(), u.DeserializeMessage, nil)
	users := new(Users)
	tao.Register(users.MessageNumber(), users.DeserializeMessage, nil)
}

func initClient(netId int64) *tao.ClientConn {
	c, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		holmes.Fatalln(err)
	}
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on close")
	})

	onMessage := tao.OnMessageOption(func(msg tao.Message, conn tao.WriteCloser) {
		fmt.Printf("%d->  %s\n", conn.(*tao.ClientConn).NetID(), msg.(*Users))
	})

	return tao.NewClientConn(netId, c, onConnect, onError, onClose, onMessage)
}

func TestClient(t *testing.T) {
	register()
	conn := initClient(1)
	for i := 0; i < 10; i++ {
		time.Sleep(60 * time.Millisecond)
		user := &User{
			Id:        i + 1,
			Name:      fmt.Sprintf("user%d", i+1),
			Password:  fmt.Sprintf("pwd%d", i+1),
			Age:       20 + uint8(i),
			Active:    true,
			CreatedAt: time.Now(),
		}
		err := conn.Write(user)
		if err != nil {
			holmes.Errorln(err)
		}
	}
	conn.Close()
}

func TestReceive(t *testing.T) {
	register()
	conn := initClient(0)
	conn.Start()
	select {}
}

func TestMoreReceive(t *testing.T) {
	register()
	for i := 0; i < 10; i++ {
		go func(i int) {
			conn := initClient(int64(i) + 1000)
			conn.Start()
		}(i)
	}
	select {}
}
