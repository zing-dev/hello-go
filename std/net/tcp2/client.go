package tcp2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func openConn() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)
	go sendMessage(conn)
	<-quitSemaphore
}

func sendMessage(conn *net.TCPConn) {
	for {
		time.Sleep(1 * time.Second)
		//b, _ := EncodeStr(string(time.Now().Year()))
		message := Message{
			Status:  true,
			Message: "OK",
		}
		bytes, _ := json.Marshal(message)
		b, _ := Encode(bytes)
		_, _ = conn.Write(b)
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := DecodeStr(reader)
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
