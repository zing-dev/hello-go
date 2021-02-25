package atomic_test

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestValueStruct(t *testing.T) {
	type A struct {
		id   int
		name string
		arr  []int
	}
	value := atomic.Value{}
	value.Store(&A{
		id:   1,
		name: "1111",
		arr:  []int{1, 2, 3},
	})
	fmt.Println(value.Load())
}
