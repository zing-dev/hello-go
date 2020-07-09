package byte

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func TestNewReader(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	log.Println(reader.Len())
	log.Println(reader.Size())
	log.Println(reader.ReadByte())
	log.Println(reader.UnreadByte())
	_, _ = reader.Seek(0, io.SeekStart)
	log.Println(reader.ReadByte())
}
