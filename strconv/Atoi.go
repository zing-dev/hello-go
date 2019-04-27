package main

import (
	"fmt"
	"strconv"
)

func main() {

	//ParseInt(s, 10, 0), converted to type int.
	fmt.Println(strconv.Atoi("1234567")) //1234567

	//const intSize = 32 << (^uint(0) >> 63)
	fmt.Println(32 << (^uint(0) >> 63))

	fmt.Println(^uint(0) >> 63) //1
	fmt.Println(^uint(0))       //18446744073709551615
	fmt.Println(uint(0))        //0
	fmt.Println(^0)             // -1
}
