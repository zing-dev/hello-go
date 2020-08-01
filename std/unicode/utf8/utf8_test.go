package utf8

import (
	"fmt"
	"math"
	"testing"
	"unicode/utf8"
)

func TestFullRune(t *testing.T) {
	t.Log(utf8.FullRune([]byte{'1', 'a'}))
	t.Log(utf8.FullRune([]byte{1, 2}))
	t.Log(utf8.FullRune([]byte{1, math.MaxInt8}))
	t.Log(utf8.FullRune([]byte{0, math.MaxInt8}))
	t.Log(utf8.FullRune([]byte{0, math.MaxInt8 + 1}))
	t.Log(utf8.FullRune([]byte{math.MaxInt8}))

	buf := []byte{228, 184, 150} // 世
	t.Log(utf8.FullRune(buf))
	t.Log(utf8.FullRune(buf[:2]))
}

func TestFullRuneInString(t *testing.T) {
	buf := "😄"
	t.Log(utf8.FullRuneInString(buf))
	t.Log(utf8.FullRuneInString(buf[:2]))
	t.Log(utf8.FullRuneInString(buf[:4]))
}

func TestDecodeRune(t *testing.T) {
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[size:]
	}

	t.Log(utf8.DecodeRune([]byte("😄")))
}

func TestDecodeRuneInString(t *testing.T) {
	t.Log(utf8.DecodeRuneInString("😄"))
}

func TestDecodeLastRune(t *testing.T) {
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[:len(b)-size]
	}

	t.Log(utf8.DecodeLastRune([]byte("😄")))
}
func TestDecodeLastRuneInString(t *testing.T) {
	t.Log(utf8.DecodeLastRuneInString("哈"))
	t.Log(utf8.DecodeLastRuneInString("😄"))
	t.Log(utf8.DecodeLastRuneInString("😄哈"))
}

func TestRuneLen(t *testing.T) {
	t.Log(utf8.RuneLen('哈'))
	t.Log(utf8.MaxRune)
	//t.Log(utf8.RuneLen('😄'))
}

func TestEncodeRune(t *testing.T) {
	r := '世'
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, r)
	t.Log(buf)
	t.Log(n)
}

func TestRuneCount(t *testing.T) {
	t.Log(utf8.RuneCount([]byte("哈")))
	t.Log(utf8.RuneCount([]byte("哈哈")))
	t.Log(utf8.RuneCount([]byte("😄")))
}

func TestRuneCountInString(t *testing.T) {
	t.Log(utf8.RuneCountInString("哈"))
	t.Log(utf8.RuneCountInString("哈哈"))
	t.Log(utf8.RuneCountInString("😄"))
}

func TestRuneStart(t *testing.T) {
	buf := []byte("a界")
	t.Log(utf8.RuneStart(buf[0]))
	t.Log(utf8.RuneStart(buf[1]))
	t.Log(utf8.RuneStart(buf[2]))
}

func TestValid(t *testing.T) {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	t.Log(utf8.Valid(valid))
	t.Log(utf8.Valid(invalid))
}

func TestValidString(t *testing.T) {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	t.Log(utf8.ValidString(valid))
	t.Log(utf8.ValidString(invalid))
}

func TestValidRune(t *testing.T) {
	valid := 'a'
	invalid := rune(0xfffffff)
	t.Log(utf8.ValidRune(valid))
	t.Log(utf8.ValidRune(invalid))
}
