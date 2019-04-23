package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1,2,3,4,5,6"
	//包含sep
	after := strings.SplitAfter(str, ",")
	fmt.Println(after)
}
