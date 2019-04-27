package main

import (
	"fmt"
	"strings"
)

func main() {

	//清空首尾空白字符
	fmt.Println(strings.TrimSpace("php is the best language over the world"))
	fmt.Println(strings.TrimSpace("hello world"))
	fmt.Println(strings.TrimSpace("hello\tworld"))
	fmt.Println(strings.TrimSpace(" hello world "))

}
