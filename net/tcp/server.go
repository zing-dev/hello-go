package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError2(err)

	fmt.Println("-------------服务器启动----------------")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError2(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//b := []byte("                     ")
		b := make([]byte, 100)
		n, err := conn.Read(b)
		checkError2(err)
		fmt.Printf("客户端：%d - %s\n", n, b)
		daytime := time.Now().String()
		_, _ = conn.Write([]byte(daytime)) // don't care about return value
		_ = conn.Close()                   // we're finished with this client
	}
}

func checkError2(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
