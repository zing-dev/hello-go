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
	Slice4()
}
func TestSlice5(t *testing.T) {
	a := make([]int, 3)
	a[100] = 1
	a[200] = 1
	a[300] = 1
	fmt.Println(a)
}
