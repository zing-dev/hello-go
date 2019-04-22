package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {

	str := "golang 大法好 😄"

	for _, val := range str {
		fmt.Println(val, strings.IndexRune(str, rune(val)))
	}
	fmt.Println("---------------------------")
	// 103 0
	// 111 1
	// 108 2
	// 97 3
	// 110 4
	// 103 0
	// 32 6
	// 22823 7
	// 27861 10
	// 22909 13
	// 32 6
	// 128516 17

	indexRune := strings.IndexRune(str, rune('😄'))
	fmt.Println(indexRune) // 17

	indexRune = strings.IndexRune(str, rune('大'))
	fmt.Println(indexRune) // 7
	indexRune = strings.IndexRune(str, rune('法'))
	fmt.Println(indexRune) // 10

	fmt.Println(len(str))      // 21
	fmt.Println(len("golang")) // 6
	fmt.Println(len("大法好"))    // 9
	fmt.Println(len("😄"))      // 4
	fmt.Println(byte('s'))
	fmt.Println(unsafe.Sizeof('😄')) // 4
	fmt.Println(unsafe.Sizeof('大')) // 4
}
