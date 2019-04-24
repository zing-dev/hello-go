package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dirName := "./io/ioutil"

	infos, e := ioutil.ReadDir(dirName)

	if e != nil {
		fmt.Println("读目录失败")
	}

	for _, info := range infos {

		fmt.Printf("当前文件%s %v 目录\n", info.Name(), info.IsDir())

	}

}
