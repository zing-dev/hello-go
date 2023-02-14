package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"

	fmt.Println(strings.ReplaceAll(str, "world", "golang"))

	fmt.Println(strings.ReplaceAll("😄🐶🐝🌶", "🐶", "😾"))
}
