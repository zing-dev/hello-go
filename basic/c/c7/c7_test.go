package c7

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	input := float32(2.33)
	output, err := Sqrt(input)
	if err != nil {
		_ = fmt.Errorf("Error: %s\n", err)
	}
	fmt.Printf("The square root of %f is %f.\n", input, output)
	Print("ABC\n")
	CallCFunc()
}
