package evio

import (
	"fmt"
	"log"
	"net"
	"testing"
)
import "github.com/tidwall/evio"

func TestClient(t *testing.T) {
	dial, err := net.Dial("tcp", "192.168.0.251:5000")
	if err != nil {
		log.Fatal(err)
	}
	dial.Write([]byte("fuck"))
	dial.Close()
}

func TestEvio(t *testing.T) {
	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = in
		fmt.Println(out)
		return
	}
	events.Closed = func(c evio.Conn, err error) (action evio.Action) {
		fmt.Println(c.AddrIndex(), "close")
		return evio.Close
	}
	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		out = []byte("hello")
		fmt.Println(c.AddrIndex(), "open")
		return
	}
	if err := evio.Serve(events, "tcp://192.168.0.251:5000"); err != nil {
		panic(err.Error())
	}
}
