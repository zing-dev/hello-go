package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var c unicode.SpecialCase = make([]unicode.CaseRange, 10)

	fmt.Println(strings.ToLowerSpecial(c, "Hello World")) //hello world

	//
	fmt.Println(string(c.ToLower('H'))) //h

}
