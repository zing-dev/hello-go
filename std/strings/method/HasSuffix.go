package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello"

	fmt.Println(strings.HasSuffix(str, "o"))  //true
	fmt.Println(strings.HasSuffix(str, "lo")) //true
	fmt.Println(strings.HasSuffix(str, ""))   //true
	fmt.Println(strings.HasSuffix(str, " "))  //false

}
