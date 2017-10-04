package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		fout *os.File
		err error
	)
	userFile := "test.txt" //文件路径
	fout, err = os.Open(userFile) //根据路径创建File的内存地址
	defer fout.Close()               //延迟关闭资源
	if err != nil {
		//fmt.Println(userFile, err)
		if os.IsNotExist(err){
			fmt.Println("file not exist,create .....")
			os.NewFile(fout.Fd(),userFile)
		}else {
			os.Remove(userFile)
		}
	}
	//循环写入数据到文件
	for i := 0; i < 10; i++ {
		fout.WriteString("Hello world!\r\n") //写入字符串
		fout.Write([]byte("abcd!\r\n"))      //强转成byte slice后再写入
	}
}
