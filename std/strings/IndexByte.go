package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	str := "hello world"

	indexByte := strings.IndexByte(str, 1)
	fmt.Println(indexByte) // -1

	indexByte = strings.IndexByte(str, 104)
	fmt.Println(indexByte)

	indexByte = strings.IndexByte(str, 111)
	fmt.Println(indexByte) // 4

	// constant 256 overflows byte
	//indexByte = strings.IndexByte(str, 256)
	//fmt.Println(indexByte)

	fmt.Printf("%d\n", 'h') // 104
	fmt.Printf("%d\n", 'o') // 111

	fmt.Println(math.MaxInt8)  //127
	fmt.Println(math.MaxUint8) // 255

}
