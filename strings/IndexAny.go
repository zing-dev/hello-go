package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world; hello golang"

	indexAny := strings.IndexAny(str, "l")
	fmt.Println(indexAny) //2

	indexAny = strings.IndexAny(str, "ho")
	fmt.Println(indexAny) // 0
}
