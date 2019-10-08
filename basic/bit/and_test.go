package bit

import (
	"fmt"
	"testing"
	"time"
)

//&

func TestAnd(t *testing.T) {
	//清零特定位 , 取某数中指定位
	a := 11
	b := 7
	fmt.Printf("%04b\n", a)   //1011
	fmt.Printf("%04b\n", b)   //0111
	fmt.Printf("%04b\n", a&b) //0011
	fmt.Println(a & b)        //3
}

func TestAnd2(t *testing.T) {
	a := time.Now().Unix()
	if a&1 == 1 {
		fmt.Println(a, "是奇数")
	} else if a&1 == 0 {
		fmt.Println(a, "是偶数")
	}
}

func TestAnd3(t *testing.T) {
	//用于消去a的最后一位的1
	a := 10
	fmt.Printf("%04b\n", a)       //1010
	fmt.Printf("%04b\n", a-1)     //1001
	fmt.Printf("%04b\n", a&(a-1)) //1000

	a = 7
	fmt.Printf("%04b\n", a)       //0111
	fmt.Printf("%04b\n", a+1)     //1000
	fmt.Printf("%04b\n", a&(a+1)) //0000

	a = 8
	fmt.Printf("%04b\n", a)       //1000
	fmt.Printf("%04b\n", a-1)     //0111
	fmt.Printf("%04b\n", a&(a-1)) //0000
}
