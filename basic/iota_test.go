package basic

import (
	"fmt"
	"testing"
)

const (
	a = iota      // 0
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
)
const (
	d = iota << 10  // 0
	e               // 1 << 10
	f               // 2 << 10
	g = iota        // 3
	h = iota * iota // 4 * 4
	i               // 5 * 5
	j = 111         // 111
	k               // 111
	l = e           // 1024
	m               // 1024
	n = iota        // 10
)

func TestIota(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
}
