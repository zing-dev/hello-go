package struct_test

import (
	"fmt"
	"math"
	"testing"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func TestMethod(t *testing.T) {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", area(r2))
	fmt.Println("Area of r2 is: ", r2.Area())

	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())

}
