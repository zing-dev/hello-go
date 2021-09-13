package demo2

import (
	"fmt"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"log"
	"net"
	"testing"
	"time"
)

func NewServer() *tao.Server {
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		holmes.Infoln("closing client")
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")
	})

	onMessage := tao.OnMessageOption(func(msg tao.Message, conn tao.WriteCloser) {
		holmes.Infoln("receive message")
		switch msg.(type) {
		case *User:
			log.Println("=> ", msg.(*User))
		case *Users:
			log.Println("=> ", msg.(*Users))
		}
	})

	server := tao.NewServer(onConnect, onClose, onError, onMessage)
	server.Sched(time.Second*5, func(t time.Time, closer tao.WriteCloser) {
		length := 10000
		users := make(Users, length)
		for i := 0; i < length; i++ {
			user := &User{
				Id:        i + 1,
				Name:      fmt.Sprintf("user%d", i+1),
				Password:  fmt.Sprintf("pwd%d", i+1),
				Age:       20 + uint8(i),
				Active:    true,
				CreatedAt: time.Now().Add(time.Duration(i) * time.Millisecond),
			}
			users[i] = user
		}
		err := closer.Write(&users)
		if err != nil {
			holmes.Errorln(err)
		}
	})
	return server
}

func TestServer(t *testing.T) {
	defer holmes.Start().Stop()
	register()
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		holmes.Fatalf("listen error %v", err)
	}
	err = NewServer().Start(l)
	if err != nil {
		t.Fatal(err)
	}
}
