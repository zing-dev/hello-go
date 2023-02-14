package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var c unicode.SpecialCase
	fmt.Println(strings.ToTitleSpecial(c, "hello world")) //HELLO WORLD

}
