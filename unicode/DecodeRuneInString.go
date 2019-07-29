package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	r, size := utf8.DecodeRuneInString("哈")
	fmt.Println(r, size)

	fmt.Println(utf8.DecodeRuneInString("a"))
	fmt.Println(utf8.DecodeRuneInString("1"))
	fmt.Println(utf8.DecodeRuneInString(""))
	fmt.Println(utf8.DecodeRuneInString(" "))

	fmt.Println(utf8.DecodeRuneInString("は"))

	str := "张荣响"
	r2, i := utf8.DecodeRuneInString(str)
	fmt.Println(r2)
	fmt.Println(str[i:])

	fmt.Println("1"[0] > +1)

}
