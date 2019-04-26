package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var c unicode.SpecialCase
	fmt.Println(strings.ToUpperSpecial(c, "hello world,哈哈")) //HELLO WORLD,哈哈

}
