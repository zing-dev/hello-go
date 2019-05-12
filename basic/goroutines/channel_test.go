package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)

		go func() {
			for i := 0; i < 10; i++ {
				a <- i
			}
			close(a)
		}()
		go func() {
			for i := range a {
				fmt.Println(i)
			}
		}()
		for i := range a {
			fmt.Println("i :", i)
		}
	}
}
