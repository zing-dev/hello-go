package test

import "fmt"

func ExampleHelloWorld() {

	fmt.Println("hello world")
	// output:
	// hello world
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
func ExamplePerm() {
	for _, value := range [...]int{0, 1, 2, 3} {
		fmt.Println(value)
	}
	// Unordered output:
	// 2
	// 1
	// 3
	// 0
}
