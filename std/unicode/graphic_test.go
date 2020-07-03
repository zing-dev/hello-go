package unicode

import (
	"log"
	"testing"
	"unicode"
)

func TestConst(t *testing.T) {
	log.Println(unicode.MaxRune)         //1114111
	log.Println(unicode.ReplacementChar) //65533
	log.Println(unicode.MaxASCII)        // 127
	log.Println(unicode.MaxLatin1)       //255
}

//letters, marks, numbers, punctuation, symbols,
//and spaces
func TestIsGraphic(t *testing.T) {
	log.Println(unicode.IsGraphic(' '))
	log.Println(unicode.IsGraphic('\a'))
	log.Println(unicode.IsGraphic('\t'))
	log.Println(unicode.IsGraphic('0'))
	log.Println(unicode.IsGraphic(0))
}

func TestIsPrint(t *testing.T) {
	log.Println(unicode.IsPrint(' '))
	log.Println(unicode.IsPrint('a'))
}

func TestIsOneOf(t *testing.T) {
	log.Println(unicode.IsOneOf(
		[]*unicode.RangeTable{unicode.ASCII_Hex_Digit},
		'1'),
	)
}

func TestIsControl(t *testing.T) {
	log.Println(unicode.IsControl(' '))
	log.Println(unicode.IsControl('\a'))
	log.Println(unicode.IsControl('\b'))
}

func TestIsLetter(t *testing.T) {
	log.Println(unicode.IsLetter('1'))
	log.Println(unicode.IsLetter('a'))
}

func TestIsMark(t *testing.T) {
	log.Println(unicode.IsMark('1'))
}
