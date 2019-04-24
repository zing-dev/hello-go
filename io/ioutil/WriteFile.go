package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	writeFileName := "./io/ioutil/test.txt"
	readFileName := "./io/ioutil/ReadFile.go"

	bytes, e := ioutil.ReadFile(readFileName)

	if e != nil {
		fmt.Printf("读文件%s失败：%s", readFileName, e)
	}

	e = ioutil.WriteFile(writeFileName, bytes, os.FileMode(0777))

	if e != nil {
		fmt.Printf("写文件%s失败：%s", writeFileName, e)
	}
	fmt.Println("写文件成功")
}
