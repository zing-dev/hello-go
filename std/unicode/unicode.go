package main

import (
	"fmt"
	"unicode"
)

func main() {
	//func IsControl(r rune) bool  // 是否控制字符
	fmt.Println(unicode.IsControl(1))    //true
	fmt.Println(unicode.IsControl(' '))  //false
	fmt.Println(unicode.IsControl('\a')) //true
	fmt.Println(unicode.IsControl('\t')) //true
	fmt.Println(unicode.IsControl('\\')) //false

	for i := 0; i <= 128; i++ {
		fmt.Printf("%3d "+
			"%3x"+
			" %7c"+
			" IsControl %v"+
			" IsDigit %v"+
			" IsGraphic %v"+
			" IsLetter %v"+
			" IsLower %v"+
			" IsMark %v"+
			" IsNumber %v"+
			" IsPrint %v"+
			" IsPunct %v"+
			" IsSpace %v"+
			" IsSymbol %v"+
			" IsTitle %v"+
			" IsUpper %v"+
			"\n",
			i,
			i,
			i,
			unicode.IsControl(rune(i)),
			unicode.IsDigit(rune(i)),
			unicode.IsGraphic(rune(i)),
			unicode.IsLetter(rune(i)),
			unicode.IsLower(rune(i)),
			unicode.IsMark(rune(i)),
			unicode.IsNumber(rune(i)),
			unicode.IsPrint(rune(i)),
			unicode.IsPunct(rune(i)),
			unicode.IsSpace(rune(i)),
			unicode.IsSymbol(rune(i)),
			unicode.IsTitle(rune(i)),
			unicode.IsUpper(rune(i)),
		)
	}

	//func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 1-9
	//func IsGraphic(r rune) bool // 是否图形字符
	//func IsLetter(r rune) bool // 是否字母
	//func IsLower(r rune) bool // 是否小写字符
	//func IsMark(r rune) bool // 是否符号字符
	//func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
	//func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个

	//func IsPrint(r rune) bool // 是否可打印字符
	//func IsPunct(r rune) bool // 是否标点符号
	//func IsSpace(r rune) bool // 是否空格
	//func IsSymbol(r rune) bool // 是否符号字符
	//func IsTitle(r rune) bool // 是否 title case
	//func IsUpper(r rune) bool // 是否大写字符
}
