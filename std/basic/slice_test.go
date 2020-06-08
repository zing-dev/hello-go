package basic

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	Slice1()
}

func TestSlice2(t *testing.T) {
	Slice2()
}

func TestSlice3(t *testing.T) {
	Slice3()
}

func TestSlice4(t *testing.T) {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)
}
