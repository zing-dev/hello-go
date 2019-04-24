package main

import (
	"fmt"
	"strconv"
)

func main() {

	f := "123.4567"

	fmt.Println(strconv.ParseFloat(f, 32))
	fmt.Println(strconv.ParseFloat(f, 64))
}
