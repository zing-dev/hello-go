package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToTitle("hello world")) //HELLO WORLD

	fmt.Println(strings.ToTitle("HELLO world")) //HELLO WORLD

	fmt.Println(strings.ToTitle("ToTitle RETURNS a copy of the strings with all Unicode letters mapped to their title case."))
}
