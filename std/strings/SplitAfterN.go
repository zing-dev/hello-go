package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1,2,3,4,5,6,7"

	//genSplit(s, sep, len(sep), n)
	//包含sep
	n := strings.SplitAfterN(str, ",", 3)
	fmt.Println(n) //[1, 2, 3,4,5,6,7]
}
