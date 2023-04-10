package basic

import (
	"fmt"
	"log"
)

func Slice1() {
	s1 := []string{"a,b,c,d"}
	for _, v := range s1 {
		v = v + v
	}
	//[a,b,c,d]
	log.Println(s1)

	s2 := &[]string{"a,b,c,d"}
	for _, v := range *s2 {
		v = v + v
	}
	//[a,b,c,d]
	log.Println(s1)

}

func Slice2() {
	c1 := []*Coder{&coder}
	for _, v := range c1 {
		v.Name = "hello kitty"
	}
	log.Println(c1[0].Name)

	c2 := []interface{}{&coder}
	for _, v := range c2 {
		v.(*Coder).Name = "hello zing"
	}
	log.Println(c2[0].(*Coder).Name)
}

type F func(str string)
type Slice struct {
	name string
	f    F
}

func (s Slice) run() {
	s.f(s.name)
}

func Slice3() {
	s1 := []Slice{{name: "zing"}, {name: "golang"}}
	for _, v := range s1 {
		v.f = func(str string) {
			log.Println(str)
		}
	}
	//invalid memory address or nil pointer dereference [recovered]
	//s1[0].run()

	s2 := []*Slice{{name: "zing"}, {name: "golang"}}
	for _, v := range s2 {
		v.f = func(str string) {
			log.Println(str)
		}
	}
	s2[0].run()
}

func Slice4() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := len(arr)
	log.Println(arr)
	start := length/2 - 2
	end := length/2 + 2
	log.Println(start, end)
	log.Println(arr[start:end])

	a := make([]int, 10)
	a[0] = 1
	a[1] = 2
	log.Println(a)
}

type S struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func RangeSlice() {
	s := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, b := range s {
		if b > 5 {
			s = append(s, b)
		}
		if b < 5 {
			s = append(s[:i], s[i+1:]...)
		}
	}
	fmt.Println(s) //[2 4 5 6 7 8 9 6 7 8 9 6 7]
}

type fn = func(slice []int, elem int) []int

func remove() []fn {
	return []fn{
		func(slice []int, elem int) []int {
			var result []int
			for _, x := range slice {
				if x != elem {
					result = append(result, x)
				}
			}
			return result
		},
		func(slice []int, elem int) []int {
			for i, x := range slice {
				if x == elem {
					if i < len(slice)-1 {
						copy(slice[i:], slice[i+1:])
					}
					slice = slice[:len(slice)-1]
					break
				}
			}
			return slice
		},
		func(slice []int, elem int) []int {
			i := 0
			for _, x := range slice {
				if x != elem {
					slice[i] = x
					i++
				}
			}
			return slice[:i]
		},
		func(slice []int, elem int) []int {
			seen := make(map[int]bool)
			var result []int
			for _, x := range slice {
				if x != elem && !seen[x] {
					seen[x] = true
					result = append(result, x)
				}
			}
			return result
		},
		func(slice []int, elem int) []int {
			var result []int
			for i, x := range slice {
				if x == elem {
					continue
				}
				if i < len(slice)-1 && i != len(result) {
					result = append(result, x)
				} else {
					result[i] = x
				}
			}
			return result
		},
	}
}
