//
// Created by zhangrongxiang on 2017/10/7 22:30
// File reader
//
package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		fileName string = "test.txt"
	)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer file.Close()    //关闭文件
	buf := make([]byte,1024)
	for{
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {

		}
		fmt.Printf("Cost time %v\n",string(buf))
	}

}
