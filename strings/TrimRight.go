package main

import (
	"fmt"
	"strings"
)

func main() {

	// TrimRight returns a slice of the string s, with all trailing
	// Unicode code points contained in cutset removed.
	fmt.Println(strings.TrimRight("hello world", "d"))    //hello worl
	fmt.Println(strings.TrimRight("hello world", "odlr")) //hello w

}
