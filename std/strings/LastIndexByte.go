package main

import (
	"fmt"
	"strings"
)

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func MyLastIndexByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func main() {
	str := "123面向接口编程😄"

	//lastIndexByte := MyLastIndexByte(str, 25509)//constant 25509 overflows byte
	//fmt.Println(lastIndexByte)
	lastIndexByte := MyLastIndexByte(str, 49)
	fmt.Println(lastIndexByte) // 0

	lastIndexByte = MyLastIndexByte(str, '2')
	fmt.Println(lastIndexByte) // 1

	lastIndexByte = strings.LastIndexByte(str, '2')
	fmt.Println(lastIndexByte) // 1

	fmt.Printf("%d-%d\n", '接', '口')
	fmt.Printf("%d-%d\n", '1', '2')
}
