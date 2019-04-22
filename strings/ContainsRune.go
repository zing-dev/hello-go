package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "1234567890"

	b := strings.ContainsRune(str, '1')
	fmt.Println(b)

	b = strings.ContainsRune(str, 49)
	fmt.Println(b)

	b = strings.ContainsRune("张", 24352)
	fmt.Println(b)

	b = strings.ContainsRune("😄👌", 128516)
	fmt.Println(b)

	fmt.Printf("%d\n", '1')
	fmt.Printf("%c\n", 49)
	fmt.Printf("%d\n", 'a')
	fmt.Printf("%c\n", 97)
	fmt.Printf("%d\n", '张')
	fmt.Printf("%d\n", '😄')
}
