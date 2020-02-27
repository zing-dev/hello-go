package main

import (
	"fmt"
	"strings"
)

// SplitN
// 用 sep 分割 string，分成 n 块,如果分到字符串不包含sep直接break
func main() {
	str := "hello world"

	n := strings.SplitN(str, " ", 2)
	fmt.Println(n)
	fmt.Println(n[0])
	fmt.Println(n[1])
	fmt.Println(len(n)) // 2

	fmt.Println(strings.SplitN("哈哈😄", "😄", 3))
	fmt.Println(strings.SplitN("哈哈😄", "😄", 3)[0])
	fmt.Println(strings.SplitN("哈哈😄", "😄", 3)[1])
	fmt.Println(len(strings.SplitN("哈哈😄", "😄", 3))) //2
	fmt.Println(cap(strings.SplitN("哈哈😄", "😄", 3))) //3

	fmt.Println(strings.SplitN("11,22,33,44,55,66", ",", 3))
	//[11 22 33,44,55,66]

}
