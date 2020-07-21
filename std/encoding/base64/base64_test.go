package base64

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}

func TestEncodeToString(t *testing.T) {
	data := []byte("any + old & data") //16 -> 18 -> 24
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
	fmt.Println(len(str))
}

func TestDecodeString(t *testing.T) {
	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
	fmt.Println(len(data))
}

func TestNewEncoder(t *testing.T) {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	_, _ = encoder.Write(input)
	_ = encoder.Close()
}
