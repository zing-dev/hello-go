package os

import (
	"fmt"
	"io"
	"log"
	"os"
)

func f() {
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

func f2() {
	name := "text.txt"
	_, err := os.Stat(name)
	if err != nil {
		//stat text.txt: no such file or directory
		fmt.Println(err)
	} else {
		err := os.Remove(name)
		if err != nil {
			fmt.Println("Remove", err)
		}
	}
	_, err = os.OpenFile(name, os.O_RDONLY, 777)
	if err != nil {
		//O_RDONLY open text.txt: no such file or directory
		fmt.Println("O_RDONLY", err)
	}
	_, err = os.OpenFile(name, os.O_WRONLY, 777)
	if err != nil {
		//O_WRONLY open text.txt: no such file or directory
		fmt.Println("O_WRONLY", err)
	}

	_, err = os.OpenFile(name, os.O_RDONLY|os.O_WRONLY, 777)
	if err != nil {
		//os.O_RDONLY|os.O_WRONLY open text.txt: no such file or directory
		fmt.Println("os.O_RDONLY|os.O_WRONLY", err)
	}
	_, err = os.OpenFile(name, os.O_RDWR, 777)
	if err != nil {
		//os.O_RDWR open text.txt: no such file or directory
		fmt.Println("os.O_RDWR", err)
	}
	_, err = os.OpenFile(name, os.O_APPEND, 777)
	if err != nil {
		//os.O_APPEND open text.txt: no such file or directory
		fmt.Println("os.O_APPEND", err)
	}
	_, err = os.OpenFile(name, os.O_EXCL, 777)
	if err != nil {
		//os.O_EXCL open text.txt: no such file or directory
		fmt.Println("os.O_EXCL", err)
	}
	_, err = os.OpenFile(name, os.O_TRUNC, 777)
	if err != nil {
		//os.O_TRUNC open text.txt: no such file or directory
		fmt.Println("os.O_TRUNC", err)
	}
	_, err = os.OpenFile(name, os.O_CREATE, 777)
	if err != nil {
		fmt.Println("os.O_TRUNC", err)
	}
}

const ConfigFile = "config.json"

func f3() {
	f, err := os.Open(ConfigFile)
	if err != nil && os.IsNotExist(err) {
		if f, err = os.Create(ConfigFile); err != nil {
			log.Println(err)
		}
	}
	_, err = f.WriteString("hello")
	if err != nil {
		//write config.json: Access is denied.
		log.Println(err)
	}
}
func f4() {
	f, err := os.OpenFile(ConfigFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString("hello")
	if err != nil {
		log.Println(err)
	}
}
