package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {

	reader := strings.NewReader("hello world")

	fmt.Println(reader.Len())  //11
	fmt.Println(reader.Size()) //11

	data := make([]byte, 10)
	fmt.Println(reader.Read(data))

	fmt.Println(string(data))

	fmt.Println(reader) //&{hello world 10 -1}

	fmt.Println(reader.Len()) //1

	fmt.Println(reader.Read(data))
	fmt.Println(string(data)) //dello worl

	_, _ = reader.Seek(0, io.SeekStart)

	fmt.Println(reader.Read(data))
	fmt.Println(string(data)) //hello worl

	//data[reader.Size()] = 'd'

	fmt.Println(len(data)) //10
	fmt.Println(cap(data)) //30
	data = append(data, 'd')
	fmt.Println(string(data)) //hello world
	fmt.Println(len(data))    //11
	fmt.Println(cap(data))    //32

	reader.Reset(string(data))

	buff := make([]byte, 10)
	buffer := bytes.NewBuffer(buff)
	_, _ = reader.WriteTo(buffer)

	fmt.Println(string(buff)) //空
	fmt.Println(buffer)       //          hello world
	fmt.Println(buffer.Len()) //21

	fmt.Println(reader.Len()) //0
	_ = reader.UnreadByte()
	fmt.Println(reader.Len()) //1
	_ = reader.UnreadByte()
	fmt.Println(reader.Len()) //2

	_ = reader.UnreadByte()
	fmt.Println(reader.Len()) //3

}
