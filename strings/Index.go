package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"

	fmt.Println(strings.Index(str, "e")) //1

	fmt.Println(strings.Index(str, "l")) //2

	fmt.Println(strings.Index(str, "")) //0

	fmt.Println(len(str))                    //11
	fmt.Println(strings.Index(str, str))     //0
	fmt.Println(strings.Index(str, str+" ")) //-1

}
