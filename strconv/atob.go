package main

import (
	"fmt"
	"strconv"
)

func MyParseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false, fmt.Errorf("解析%s错误", str)
}

func main() {

	fmt.Println(MyParseBool("false"))
	fmt.Println(MyParseBool("false1"))

	b, e := MyParseBool("0")
	if e != nil {
		fmt.Println(b)
	}
	fmt.Println(b) //false

	fmt.Println(strconv.ParseBool("1")) //true

	fmt.Println(strconv.FormatBool(true))  //true
	fmt.Println(strconv.FormatBool(false)) // false

	dst := make([]byte, 0)
	dst = strconv.AppendBool(dst, true)
	fmt.Println(dst)         //[116 114 117 101]
	fmt.Println(string(dst)) //true
	fmt.Printf("%s\n", dst)  //true

	dst = strconv.AppendBool(dst, false)
	fmt.Println(string(dst)) //truefalse

}
