package basic

import (
	"fmt"
	"math"
	"testing"
)

func TestFmt(t *testing.T) {
	fmt.Printf("%d\n", 1)
	fmt.Printf("%.d\n", 1)
	fmt.Printf("%.2d\n", 1)
	fmt.Printf("%.2d\n", 11)
	fmt.Printf("%.2d\n", 111)
	fmt.Printf("%.2d%.2d\n", 111, 222)
	fmt.Printf("%.2b\n", 111)
	fmt.Printf("%.2x\n", 111)
	fmt.Printf("%v\n", []byte{1, 11, 111, 222})
}

type Point struct {
	x float64
	y float64
}

func TestPoint(t *testing.T) {
	//288 257 568 256 367 255
	fmt.Println(calculation(Point{300, 250}, Point{600, 250}, Point{367, 250}))
	fmt.Println(calculation(Point{288, 257}, Point{568, 256}, Point{367, 255}))

	fmt.Println(calculation(Point{1, 0.5}, Point{4, 4}, Point{1, 1}))
}
func calculation(p1, p2, p3 Point) float64 {
	var a float64
	var b float64
	a = (p1.y - p2.y) / (p1.x - p2.x)
	b = p1.y - a*p1.x
	l1 := math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
	h1 := math.Abs(p3.y - a*p3.x - b)
	w1 := math.Abs(p3.x - (p3.y-b)/a)
	if math.IsNaN(h1) {
		h1 = 0
	}
	if math.IsNaN(w1) {
		w1 = 0
	}
	return h1 * w1 / l1
}

func TestSprintf(t *testing.T) {
	fmt.Println(fmt.Sprintf("%6s", "hi"))
	fmt.Println(fmt.Sprintf("%-6s", "hi"))
	fmt.Println(fmt.Sprintf("%10s", "hi"))
	fmt.Println(fmt.Sprintf("%-10s", "hi"))
	fmt.Println(fmt.Sprintf("%-10s = v", "k"))
	fmt.Println(fmt.Sprintf("%-10s = v2", "k2"))
	fmt.Println(fmt.Sprintf("%-10s = %-2s", "k3", "v3"))
	fmt.Println(fmt.Sprintf("%-10s = %-5s", "k3", "v3"))
	fmt.Println(fmt.Sprintf(fmt.Sprintf("%%-%ds = %%-%ds", 10, 5), "hi", "wocao"))
}
