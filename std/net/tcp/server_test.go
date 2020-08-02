package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	service := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------服务器启动----------------")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		for {
			b := make([]byte, 100)
			n, err := conn.Read(b)
			if b[0] == 'a' || err == io.EOF {
				log.Println("客户端退出")
				break
			}
			fmt.Printf("客户端%s ：%d - %s\n", conn.RemoteAddr().String(), n, b)
			daytime := time.Now().Format("2006-01-02 15:04:05")
			n, err = conn.Write([]byte(daytime)) // don't care about return value
			if err != nil {
				break
			}
			fmt.Printf("send %d:%s \n", n, daytime)
		}
		//_ = conn.Close() // we're finished with this client
	}
}
