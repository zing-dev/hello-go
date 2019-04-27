package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(strconv.ParseUint("12", 10, 0)) //12

	fmt.Println(strconv.ParseUint("12", 8, 0)) // 10

	fmt.Println(strconv.ParseUint("11111", 2, 16)) //31

}
