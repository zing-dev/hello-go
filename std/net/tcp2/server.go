package tcp2

import (
	"bufio"
	"fmt"
	"net"
)

var ConnMap map[string]*net.TCPConn

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		_ = conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := DecodeStr(reader)
		if err != nil {
			return
		}
		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))

		b, err := EncodeStr(conn.RemoteAddr().String() + ":" + string(message))
		if err != nil {
			continue
		}
		_, _ = conn.Write(b)

	}
}

func boradcastMessage(message string) {
	for _, conn := range ConnMap {
		b, err := EncodeStr(message)
		if err != nil {
			continue
		}
		_, _ = conn.Write(b)
	}
}
