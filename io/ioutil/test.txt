package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	bytes, e := ioutil.ReadFile(`C:\Users\zhang\workspace\go\learn-go\io\ioutil\ReadFile.go`)

	if e != nil {
		fmt.Println("读文件失败：", e)
	}

	fmt.Println(bytes)

	fmt.Printf("%s", bytes)
}
