//
// Created by zhangrongxiang on 2017/9/30 10:57
// File recursion2
//
package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 20; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
