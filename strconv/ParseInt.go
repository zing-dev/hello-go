package main

import (
	"fmt"
	"strconv"
)

func main() {

	// If base == 0, the base is implied by the string's prefix:
	// base 16 for "0x", base 8 for "0", and base 10 otherwise.
	// For bases 1, below 0 or above 36 an error is returned.
	fmt.Println(strconv.ParseInt("0xf", 0, 8)) //15
	fmt.Println(strconv.ParseInt("20", 0, 8))  //20
	fmt.Println(strconv.ParseInt("20", 10, 8)) //20
	//fmt.Println(strconv.ParseInt("0xf",10,8))

	fmt.Println(strconv.ParseInt("1010", 2, 8)) //10

	fmt.Println(strconv.ParseInt("1010", 3, 8)) //30

	fmt.Println(strconv.ParseInt("-1010", 8, 16)) //-520

	fmt.Println(2 ^ 0) // 1

}
