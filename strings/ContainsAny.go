package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world, hello golang"
	c := strings.ContainsAny(str, "h")
	fmt.Println(c)

	c = strings.ContainsAny(str, " ")
	fmt.Println(c)

	c = strings.ContainsAny(str, "world")
	fmt.Println(c)
}
