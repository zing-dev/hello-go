package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("hello", "Hello")) //true

	fmt.Println(strings.EqualFold("😄", "😄")) //true

	fmt.Println(strings.EqualFold("😄1", "😄2")) //false

	fmt.Println(strings.EqualFold("あ", "ア")) //false

	fmt.Println(strings.EqualFold("い", "イ")) //false
}
