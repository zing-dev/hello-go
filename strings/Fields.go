package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"

	//space characters,
	fields := strings.Fields(str)
	fmt.Println(len(fields))
	for k, v := range fields {
		fmt.Println(k, v)
	}
}
