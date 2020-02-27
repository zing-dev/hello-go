package main

import (
	"fmt"
	"strings"
)

func main() {

	c := strings.Count("hello golang", "l")
	fmt.Println(c)

	str := "golang是世界上最好的语言哈哈哈哈哈;😄😄"
	c = strings.Count(str, "哈")
	fmt.Println(c)

	c = strings.Count(str, "😄")
	fmt.Println(c)

}
