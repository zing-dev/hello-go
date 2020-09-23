package bit

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	a := 11
	b := 7
	fmt.Printf("%08b\n", a)   //00001011
	fmt.Printf("%08b\n", ^a)  //-0001100
	fmt.Printf("%08b\n", b)   //00000111
	fmt.Printf("%08b\n", ^b)  //-0001000
	fmt.Printf("%08b\n", -7)  //-0000111
	fmt.Printf("%08b\n", ^-7) //00000110
	fmt.Printf("%d\n", ^-7)   //6

	fmt.Printf("%08b\n", a^b) //00001100
}

func TestXor2(t *testing.T) {
	//二进制数在内存中以补码的形式存储。
	//按位取反：二进制每一位取反，0变1，1变0。
	//~9的计算步骤：
	//转二进制：0 1001
	fmt.Printf("%08b\n", 9) //00001001
	//计算补码：0 1001
	//按位取反：1 0110

	//转为原码：
	//按位取反：1 1001
	//末位加一：1 1010
	//符号位为1是负数，即-10
	fmt.Printf("%08b\n", ^9) //-0001010
}
func TestXor3(t *testing.T) {
	//使特定位的值取反 (mask中特定位置1，其它位为0 s=s^mask)
	a := 11
	mask := 8
	fmt.Printf("%08b\n", a)      //00001011
	fmt.Printf("%08b\n", mask)   //00001000
	fmt.Printf("%08b\n", mask^a) //00000011
}

func TestXor4(t *testing.T) {
	a := 1
	b := 2

	c := a ^ b
	fmt.Println(c)
	a = c ^ a
	fmt.Println(a) //2
	b = c ^ b
	fmt.Println(b) //1

}

func TestXor5(t *testing.T) {
	a := 1
	b := 2

	fmt.Println(a ^ b)     //3
	fmt.Println(a ^ b ^ a) //2
	fmt.Println(a ^ b ^ b) //1

	a ^= b ^ a
	fmt.Println(a) //2
	b ^= b ^ a
	fmt.Println(b) //0

	a ^= b
	fmt.Println(a ^ b)     //2
	fmt.Println(a ^ b ^ a) //1
	fmt.Println(1 ^ 2 ^ 1) //1

} //

func TestXor6(t *testing.T) {
	a := 1
	b := 2
	a ^= b
	fmt.Println(a) //3
	b ^= a
	fmt.Println(b) //1
	a ^= b
	fmt.Println(a) //2
}

func TestXor7(t *testing.T) {
	u := binary.BigEndian.Uint32([]byte{0, 0, 0, 254})
	fmt.Println((u & 1 << 0) >> 0)

	//11000000
	u = binary.BigEndian.Uint32([]byte{0, 0, 0, 192})
	fmt.Println(u & (1 << 6))
	fmt.Println(u & (1 << 6) >> 6)

	fmt.Println(fmt.Sprintf("%8b", 192))  //11000000
	fmt.Println(fmt.Sprintf("%8b", 1<<5)) //
	fmt.Println(fmt.Sprintf("%b", 192&(1<<5)))
	fmt.Println(fmt.Sprintf("%b", 192^(1<<7)))
	fmt.Println(fmt.Sprintf("%b", (192^(1<<7))>>6))

	fmt.Println(fmt.Sprintf("%b", 254))
	fmt.Println(fmt.Sprintf("%b", 254&1))
}
