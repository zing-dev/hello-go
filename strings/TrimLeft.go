package main

import (
	"fmt"
	"strings"
)

func main() {

	cfgs := "mongodb://off"
	cfgs = strings.TrimLeft(cfgs, "mongodb://")
	fmt.Printf("cfgs:%v\n", cfgs)
	//output ==> cfgs:ff

	// TrimLeft returns a slice of the string s with all leading
	// Unicode code points contained in cutset removed.
	fmt.Println(strings.TrimLeft("hello", "h l o e")) //空
}
