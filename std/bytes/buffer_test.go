package byte

import (
	"bytes"
	"log"
	"testing"
)

var buffer = bytes.NewBufferString("hello world")

func TestNewBufferString(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	log.Println(buf.Bytes())
	log.Println(buf.Cap())
	log.Println(buf.Len())
	log.Println(buf.String())
}

func TestNewBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hello world"))
	log.Println(buf.String())
}

func TestTruncate(t *testing.T) {
	buffer.Truncate(2)
	log.Println(buffer.String())
}

func TestGrow(t *testing.T) {
	log.Println(buffer.Len())
	log.Println(buffer.Cap())
	buffer.Grow(buffer.Cap())
	log.Println(buffer.Len())
	log.Println(buffer.Cap())
}
