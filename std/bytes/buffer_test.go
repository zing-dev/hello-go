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

func TestWrite(t *testing.T) {
	buffer.Reset()
	buffer.Write([]byte("hello golang"))
	log.Println(string(buffer.Bytes()))
}

func TestWriteString(t *testing.T) {
	buffer.Reset()
	buffer.WriteString("golang")
	log.Println(string(buffer.Bytes()))
}

func TestReadFrom(t *testing.T) {
	buf := bytes.NewBufferString("hello golang")
	buffer.Reset()
	n, err := buffer.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buffer.String())
}

func TestWriteTo(t *testing.T) {
	buf := bytes.Buffer{}
	n, err := buffer.WriteTo(&buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buf.String())
}

func TestWriteByte(t *testing.T) {
	buffer.Reset()
	buffer.WriteByte('a')
	buffer.WriteByte('b')
	buffer.WriteByte('c')
	log.Println(buffer.String())
}

func TestWriteRune(t *testing.T) {
	buffer.Reset()
	buffer.WriteRune('哈')
	buffer.WriteRune('喽')
	log.Println(buffer.String())
}

func TestRead(t *testing.T) {
	buf := make([]byte, buffer.Len())
	n, err := buffer.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buf)
	log.Println(buffer.Len())
}
