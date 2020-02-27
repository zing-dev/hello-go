package main

import (
	"fmt"
	"strings"
)

func main() {

	//清除指定的首尾 字符串
	fmt.Println(strings.Trim("\nhello boys and girls\n", "\n"))
}
