package tcp2

import (
	"encoding/json"
	"fmt"
	"net"
	"testing"
)

func TestName(t *testing.T) {
	message := Message{
		Status:  true,
		Message: "OK",
	}

	bytes, _ := json.Marshal(message)
	fmt.Println(len(bytes))
}
func TestClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		go openConn()
	}
	var msg string
	_, _ = fmt.Scanln(&msg)
	<-quitSemaphore
}

func TestServer(t *testing.T) {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		go tcpPipe(tcpConn)
	}
}
