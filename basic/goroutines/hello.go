package main

import "fmt"

func say(name string) {
	fmt.Println("hello ", name)
}

func main() {

	go say("goroutine")
	go say("world")
	go say("golang")

	fmt.Println("main")
}
