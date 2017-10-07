//
// Created by zhangrongxiang on 2017/10/7 22:07
// File write
//
package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	var (
		fout *os.File
		err error
		fileName string = "test.txt"
	)

	fout,err = os.OpenFile(fileName,os.O_RDWR | os.O_APPEND,0777)
	if err != nil {
		fmt.Println("openFile error")
		return
	}
	defer fout.Close()
	i := 1
	for i < 20{
		fout.WriteString( "today is beautiful     " + time.Now().String() +"\n")
		time.Sleep(time.Second)
		i++
	}

}