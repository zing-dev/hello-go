package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := append(a, 3)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b)) // 2 4
	c := append(b, 4)
	d := append(b, 5)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b)) // 2 4
	fmt.Printf("cap(c)=%v,cap(d)=%v\n", cap(c), cap(d)) // 4 4
	fmt.Println(a, b, c, d)
	fmt.Printf("%p,%p,%p,%p\n", a, b, c, d)

	e := append(d, 6, 7)
	f := append(e, 8, 9, 10)
	fmt.Printf("cap(e)=%v,cap(f)=%v\n", cap(e), cap(f)) // 8 16
	fmt.Println(f)                                      //[1 2 3 5 6 7 8 9 10]

	g := append(e, 11, 12, 13, 14, 15, 16)
	fmt.Printf("cap(g)=%v,cap(g)=%v\n", cap(g), cap(g)) // 8 16

	h := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := append(h, 11, 12, 13, 14)
	fmt.Println(cap(h), cap(i)) //10 32

}
