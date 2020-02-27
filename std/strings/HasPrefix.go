package main

import (
	"fmt"
	"strings"
)

type S struct {
	string
}

func (s S) String() string {
	return s.string
}
func main() {
	str := "hello"

	fmt.Println(strings.HasPrefix(str, "h"))  //true
	fmt.Println(strings.HasPrefix(str, "he")) //true
	fmt.Println(strings.HasPrefix(str, "e"))  //false

	s := S{
		"hello world",
	}

	fmt.Println(s)
}
