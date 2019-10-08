package bit

import (
	"fmt"
	"testing"
)

func TestRightShift(t *testing.T) {
	fmt.Println(1 >> 0)
	fmt.Println(1 >> 1)
	fmt.Println(1 >> 2)

	fmt.Println(128 >> 0)
	fmt.Println(128 >> 1)
	fmt.Println(128 >> 2)

	fmt.Println(-128 >> 0)
	fmt.Println(-128 >> 1)
	fmt.Println(-128 >> 2)
}
