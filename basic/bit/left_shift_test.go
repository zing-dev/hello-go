package bit

import (
	"fmt"
	"testing"
)

func TestLeftShift(t *testing.T) {
	a := 11
	b := 7
	fmt.Printf("%08b\n", a)    //00001011
	fmt.Printf("%08b\n", a<<1) //00010110
	fmt.Printf("%08b\n", b)    //00000111
	fmt.Printf("%08b\n", b<<1) //00001110
}

func TestLeftShift2(t *testing.T) {
	fmt.Printf("%08b\n", 1<<0) //00000001
	fmt.Printf("%08b\n", 1<<1) //00000010
	fmt.Printf("%08b\n", 1<<2) //00000100
	fmt.Printf("%08b\n", 1<<3) //00001000
	fmt.Printf("%08b\n", 1<<4) //00010000

}

func TestLeftShift3(t *testing.T) {
	//将int型变量a的第k位置1，即a=a|(1<<k)
	a := 11                        //00001011
	fmt.Printf("%08b\n", 1<<4)     //00010000
	fmt.Printf("%08b\n", a|(1<<4)) //00011011

	//将int型变量a的第k位清0，即a=a&~(1<<k)    (10000 取反后为00001 )
	fmt.Printf("%d\n", 1<<3)        //8
	fmt.Printf("%08b\n", 1<<3)      //00001000
	fmt.Printf("%d\n", ^(1 << 3))   //-9
	fmt.Printf("%08b\n", ^(1 << 3)) //-0001001
	fmt.Printf("%08b\n", a&^8)      //00000011
	fmt.Printf("%08b\n", a&-9)      //00000011
	fmt.Printf("%08b\n", a&^(1<<3)) //00000011
}
