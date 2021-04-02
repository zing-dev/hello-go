package tcp

import (
	"io"
	"log"
	"net"
	"testing"
	"time"
)

const (
	Read = iota
	ReadAll
)

//网络读取数据不要用ReadAll
func TestTcpServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":1100")
	if err != nil {
		t.Fatal(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		t.Fatal(err)
	}
	switch ReadAll {
	case ReadAll:
		d, err := io.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}
		log.Println(len(d))
	case Read:
		data := make([]byte, 1024)
		for {
			n, err := conn.Read(data)
			if err != nil {
				t.Fatal(err)
			}
			log.Println(n)
		}
	}
}

func TestTcpClient(t *testing.T) {
	conn, err := net.DialTimeout("tcp", ":1100", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]byte, 1024*1024)
	n, err := conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(n)
	time.Sleep(time.Second * 10)
}
