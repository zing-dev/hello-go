package bit

import (
	"fmt"
	"testing"
)

func TestOr(t *testing.T) {
	a := 11
	b := 7
	fmt.Printf("%04b\n", a)   //1011
	fmt.Printf("%04b\n", b)   //0111
	fmt.Printf("%04b\n", a|b) //1111
	fmt.Println(a | b)        //15
}

func TestOr2(t *testing.T) {

	//把最后一位变成1
	fmt.Printf("%04b\n", 12)   //1100
	fmt.Printf("%04b\n", 12|1) //1101

	//把最后一位变成0
	fmt.Printf("%04b\n", 12)     //1100
	fmt.Printf("%04b\n", 12|1-1) //1100

}

func TestOr3(t *testing.T) {
	a := 11
	//把右起第一个0变成1
	fmt.Printf("%04b\n", a|(a+1)) //1111
	//把右边连续的0变成1,必须是偶数
	a = 128
	fmt.Printf("%08b\n", a)       //10000000
	fmt.Printf("%08b\n", a-1)     //01111111
	fmt.Printf("%08b\n", a|(a-1)) //11111111
}
