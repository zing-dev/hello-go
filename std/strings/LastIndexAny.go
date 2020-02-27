package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "123面向接口编程😄"
	lastIndexAny := strings.LastIndexAny(str, "😄")
	fmt.Println(lastIndexAny) // 3 + 3 * 6

	fmt.Println(len("😄")) //4
	fmt.Println(len("程")) //3
}
