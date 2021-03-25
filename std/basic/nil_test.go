package basic

import (
	"io"
	"log"
	"testing"
)

const relayMaxSize = 256
const relayMinSize = 4

func TestNil(t *testing.T) {
	t.Log(nil)
}

func TestArr(t *testing.T) {
	var data [relayMaxSize]byte
	n, err := io.ReadAtLeast(nil, data[:], relayMinSize)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(n)

}
