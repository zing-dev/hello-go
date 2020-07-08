package basic

import (
	"fmt"
	"log"
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

func TestSlice6(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	log.Println(arr)
	arr = append(arr[:3], arr[4:]...)
	log.Println(arr)
	arr = append(arr, 11)
	log.Println(arr)

	m := map[string][]int{}
	m["1"] = []int{1, 2, 3}
	m["1"][0] = 11
	one := m["1"]
	one[0] = 12
	one = []int{}

	m["2"] = arr
	arr = append(arr[:3], arr[9:]...)
	arr = m["2"]
	arr = []int{}

	log.Println(m)
}
