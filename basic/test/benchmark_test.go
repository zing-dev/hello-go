package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkIntn(b *testing.B) {

	for i := 0; i < b.N; i++ {
		r := rand.Intn(1000000)

		if r > 999998 {
			b.Error("err", r)
		}
	}
}
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}
