package main

import (
	"fmt"
	"strings"
)

func main() {

	str := strings.Map(func(r rune) rune {
		return r
	}, "hello world")

	fmt.Println(str) //hello world

	str = strings.Map(func(r rune) rune {
		return r - 1
	}, "ifmmp!xpsme")

	fmt.Println(str) //hello world

}
