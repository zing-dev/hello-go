package tcp

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	network := "tcp4"
	address := ":7777"
	addr, e := net.ResolveTCPAddr(network, address)
	checkError(e)

	conn, e := net.DialTCP("tcp", nil, addr)
	checkError(e)
	//_, e = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	_, e = conn.Write([]byte("hello world,hello server"))
	checkError(e)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
