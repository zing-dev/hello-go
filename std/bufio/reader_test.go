package bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

var reader = bufio.NewReader(bytes.NewBufferString("hello world!"))

func TestRead(t *testing.T) {
	data := make([]byte, 3)
	for {
		n, err := reader.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, string(data))
	}
}

func TestNewReaderSize(t *testing.T) {
	reader := bufio.NewReaderSize(bytes.NewBufferString("hello world"), 2)
	log.Println(reader.Size())
	c, err := reader.ReadByte()
	log.Println(fmt.Sprintf("%c %v", c, err))

	data := make([]byte, 2)
	n, err := reader.Read(data)
	log.Println(n, err, string(data))

	log.Println(reader.UnreadByte())

	log.Println(reader.Size())
	log.Println(reader.Buffered())
}

func TestNewReader(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBuffer([]byte("hello golang")))
	log.Println(reader.Size())
}

func TestPeek(t *testing.T) {
	reader := bufio.NewReader(bytes.NewReader([]byte("hello")))
	peek, err := reader.Peek(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(peek))
	peek, err = reader.Peek(1)
	log.Println(string(peek), err)
}

func TestDiscard(t *testing.T) {
	discard, err := reader.Discard(2)
	log.Println(discard, err)
	b, err := reader.ReadByte()
	log.Println(string(b), err)
}

func TestReadSlice(t *testing.T) {
	line, err := reader.ReadSlice(' ')
	log.Println(string(line), err)
	b, err := reader.ReadByte()
	log.Println(string(b), err)
}

func TestReadLine(t *testing.T) {
	data, _ := os.ReadFile("reader_test.go")
	reader := bufio.NewReader(bytes.NewReader(data))
	for {
		line, prefix, err := reader.ReadLine()
		fmt.Println(string(line), prefix, err)
		if err == io.EOF {
			fmt.Println("io.EOF break")
			break
		}
		if err != nil {
			fmt.Println(err, "break")
			break
		}
	}
}

func TestReadBytes(t *testing.T) {
	d, err := reader.ReadBytes(' ')
	log.Println(string(d), err)
}

func TestNewWriterSize(t *testing.T) {
	b := make([]byte, 1024)
	buf := bytes.NewBuffer(b)
	write := bufio.NewWriterSize(buf, 64)
	log.Println(write.Size())
	n, err := write.WriteString("hello world")
	log.Println(n, err)
	log.Println(string(buf.Bytes()))
}
