package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

// EOF is the error returned by Read when no more input is available.
// (Read must return EOF itself, not an error wrapping EOF,
// because callers will test for EOF using ==.)
// Functions should return EOF only to signal a graceful end of input.
// If the EOF occurs unexpectedly in a structured data stream,
// the appropriate error is either ErrUnexpectedEOF or some other error
// giving more detail.
//var EOF = errors.New("EOF")

// ErrUnexpectedEOF means that EOF was encountered in the
// middle of reading a fixed-size block or data structure.
//var ErrUnexpectedEOF = errors.New("unexpected EOF")

func TestReadFull(t *testing.T) {
	buffer := bytes.NewBufferString("hello world!")
	data := make([]byte, 3)
	for {
		n, err := io.ReadFull(buffer, data)
		fmt.Println(string(data[:n]), n, err)
		if err != nil {
			if err == io.EOF {
				log.Println("eof...", n)
			} else {
				log.Println("other err", err, n)
			}
			break
		}
		if n == 0 {
			fmt.Println("n == 0 break")
			break
		}
	}

	buffer = bytes.NewBufferString("buffer = bytes.NewBufferString()")
	dst := bytes.Buffer{}
	for {
		n, err := buffer.Read(data)
		dst.Write(data[:n])
		if err != nil {
			break
		}
	}
	data = dst.Bytes()
	fmt.Println(string(data), len(data))
}

func TestCopy(t *testing.T) {
	src := bytes.NewBufferString("hello world!")
	dst := &bytes.Buffer{}
	n, err := io.Copy(dst, src)
	log.Println(n, string(dst.Bytes()), err)

	src = bytes.NewBuffer(make([]byte, 1024*1024*1024))
	n, err = io.Copy(dst, src)
	log.Println(n, 1024*1024*1024, err)
}
