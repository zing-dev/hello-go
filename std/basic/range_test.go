package basic

import (
	"fmt"
	"testing"
)

type Str struct {
	bar string
}

//range每次都会把当前值赋值到循环变量上，而不是直接使用原变量，在例程中，我们保存了当前循环变量的地址，但是Go每次循环都会复用这一"循环变量",
//所以每次保存的地址实际上指向同一变量，最后，内容都变成了最后一个元素的值。
func TestRange(t *testing.T) {
	list := []Str{
		{"1"},
		{"2"},
		{"3"},
	}
	list0 := []*Str{
		{"1"},
		{"2"},
		{"3"},
	}
	list2 := make([]*Str, len(list))
	list3 := make([]Str, len(list))
	for i, value := range list {
		list2[i] = &value
		list3[i] = value
	}

	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2]) //&{3} &{3} &{3}
	fmt.Println(list3[0], list3[1], list3[2])

	for i, value := range list0 {
		list2[i] = value
		list3[i] = *value
	}
	fmt.Println(list0[0], list0[1], list0[2])
	fmt.Println(list2[0], list2[1], list2[2])
	fmt.Println(list3[0], list3[1], list3[2])
}
