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

func TestNext(t *testing.T) {
	log.Println(buffer.Next(1))
	log.Println(buffer.Next(2))
}

func TestReadByte(t *testing.T) {
	b, err := buffer.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b), b)
}

func TestReadRune(t *testing.T) {
	r, size, err := buffer.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r, size)

	buffer.Reset()
	buffer.WriteRune('哈')
	r, size, err = buffer.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r, size)
}

func TestUnreadRune(t *testing.T) {
	log.Println(buffer.UnreadByte())
	buffer.Reset()
	log.Println(buffer.UnreadByte())
	buffer.Next(1)
	log.Println(buffer.UnreadByte())
	buffer.WriteByte('a')
	log.Println(buffer.UnreadByte())
	_, _ = buffer.ReadByte()
	log.Println(buffer.UnreadByte())
}

func TestReadBytes(t *testing.T) {
	line, err := buffer.ReadBytes(' ')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(line, string(line))
}

func TestReadString(t *testing.T) {
	line, err := buffer.ReadString(',')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(line)
}
