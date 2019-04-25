package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./os/test.txt"

	//file, e := os.Open(fileName)
	file, e := os.OpenFile(fileName, os.O_RDWR, os.FileMode(0777))
	if e != nil {
		fmt.Println("open error:", e)
		fmt.Println("start create")

		file, e = os.Create(fileName)

		if e != nil {
			fmt.Println("create error:", e)
		}
	}
	fmt.Printf("%s 存在\n", file.Name())

	data := make([]byte, 1024)
	n, e := file.Read(data)

	if n == 0 {
		fmt.Println("这个文件还没有任何内容，开始写内容")
		i, e := file.Write([]byte("这是一段文件内容，😄\n"))
		if e != nil {
			fmt.Println("写文件失败：", i)
		}
		fmt.Println("第一次写入文件成功，长度为：", i)
	} else {
		fmt.Println(string(data))
	}

	i, e := file.WriteString("hello golang,😘\n")
	if e != nil {
		fmt.Println("写入内容失败", e)
	}

	fmt.Println("写入的字符长度为：", i)

	d := make([]byte, 10)

	_, e = file.Seek(0, io.SeekStart)

	for true {
		n2, e := file.Read(d)
		fmt.Print(string(d))
		if e != nil && e != io.EOF {
			panic(e)
		}
		if n2 == 0 {
			break
		}
	}
}
