package main

import (
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
