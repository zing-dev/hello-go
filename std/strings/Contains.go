package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"

	contains := strings.Contains(str, "ll")
	fmt.Println(contains)

	contains = strings.Contains(str, "\b")
	fmt.Println(contains)

	fmt.Println("|\b|\t|")
}
