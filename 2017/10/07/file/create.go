//
// Created by zhangrongxiang on 2017/10/7 17:25
// File create
//
package main

import (
	"os"
	"fmt"
)

func main() {
	var  (
		fout     *os.File
		err      error
		fileName string = "test.txt"
	)

	fout,err = os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err){
			fout, err = os.Create(fileName)
		}else {
			fmt.Println("fout aleraly exist...")
		}
	}else{
		fmt.Println("no error")
	}
	name := fout.Name()
	fmt.Printf("fout name is %s\n",name)
	fd := fout.Fd()
	fmt.Printf("fout fd is %d\n",fd)
	info, err := fout.Stat()
	if err != nil {
		fmt.Println("error")
	}
	if info.IsDir(){
		fmt.Println("fout is dir")
	}else {
		fmt.Println("fout is fout")
	}

	fmt.Println(info)
	err = fout.Close()
	if err != nil {
		fmt.Println(os.ErrClosed)
	}

	fmt.Println()
}
