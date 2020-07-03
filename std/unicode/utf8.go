package unicode

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {

	//fmt.Println(utf8.FullRune([]byte{'1', 'a', '张', '😄'}))
	fmt.Println(utf8.FullRune([]byte{'1', 'a'}))
	fmt.Println(utf8.FullRune([]byte{1, 2}))
	fmt.Println(utf8.FullRune([]byte{1, math.MaxInt8}))
	fmt.Println(utf8.FullRune([]byte{0, math.MaxInt8}))
	//fmt.Println(utf8.FullRune([]byte{0, '哈'}))
	//fmt.Println(utf8.FullRune([]byte{0, math.MaxInt8 + 1}))

	fmt.Println(utf8.FullRuneInString("😄"))

	fmt.Println(utf8.DecodeRune([]byte{100, 101}))
	//fmt.Println(utf8.DecodeRune([]byte{'张', '😄'}))

	fmt.Println(utf8.RuneLen(10))    //1
	fmt.Println(utf8.RuneLen(10000)) //3
	fmt.Println(utf8.RuneLen('张'))   //3
	//fmt.Println(utf8.RuneLen('😄'))   //4

	fmt.Printf("%d\n", '张')
}
