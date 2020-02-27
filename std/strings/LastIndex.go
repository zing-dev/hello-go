package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"

	index := strings.LastIndex(str, "l")
	fmt.Println(index) //9

	index = strings.LastIndex(str, " ")
	fmt.Println(index) //5

	index = strings.LastIndex("golang大法好", "大法")
	fmt.Println(index) //6
}
