package test

import (
	"fmt"
	"testing"
)

func BenchmarkHelloParallel(b *testing.B) {

	fmt.Println("hello ")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Println("world")
		}
	})
}
