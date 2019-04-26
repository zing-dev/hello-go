package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1,2,3,4,5,6,7,8,9"

	fields := strings.FieldsFunc(str, func(r rune) bool {
		return r == ','
	})
	fmt.Println(fields) //[1 2 3 4 5 6 7 8 9]

	fields = strings.FieldsFunc(str, func(r rune) bool {
		i, _ := strconv.Atoi(string(r))
		//fmt.Printf("%v %d\n", r, i)
		return i%2 == 0
	})

	fmt.Println(fields) //[1 3 5 7 9]

}
