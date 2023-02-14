package main

import (
	"fmt"
	"strings"
	"testing"
)

func NumberStrToSlice(str string) {
	for _, v := range strings.Split(str, ",") {
		println(v)
	}
}

func TestNumberStrToSlice(t *testing.T) {
	NumberStrToSlice("1,2,3,4,5")
}

func TestChar(t *testing.T) {
	fmt.Println('P')
	fmt.Println(string('P'))
	fmt.Println(string('P') == "P")
	fmt.Println(string(rune(80)))
	fmt.Println(string(byte(80)))
}
