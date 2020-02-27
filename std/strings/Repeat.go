package main

import (
	"fmt"
	"strings"
)

func main() {

	repeat := strings.Repeat("golang\t", 5)
	fmt.Println(repeat) //golang	golang	golang	golang	golang

}
