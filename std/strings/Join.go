package main

import (
	"fmt"
	"strings"
)

func main() {
	strSlice := []string{"hello", "world", "hello", "golang"}

	strJoin := strings.Join(strSlice, " ")

	fmt.Println(strJoin) //hello world hello golang
}
