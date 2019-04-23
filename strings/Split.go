package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1,2,3,4,5,6"

	split := strings.Split(str, ",")
	fmt.Println(split) //[1 2 3 4 5 6]

	split = strings.SplitN(str, ",", -1)
	fmt.Println(split) //[1 2 3 4 5 6]
}
