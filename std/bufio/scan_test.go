package bufio

import (
	"bufio"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestSpace(t *testing.T) {
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if unicode.IsSpace(r) {
			t.Fatalf("white space property disagrees: %#U should be %t", r, unicode.IsSpace(r))
		}
	}
}

var scanTests = []string{
	"",
	"a",
	"¼",
	"☹",
	"\x81",   // UTF-8 error
	"\uFFFD", // correctly encoded RuneError
	"abcdefgh",
	"abc def\n\t\tgh    ",
	"abc¼☹\x81\uFFFD日本語\x82abc",
}

func TestScanByte(t *testing.T) {
	for n, test := range scanTests {
		buf := strings.NewReader(test)
		s := bufio.NewScanner(buf)
		s.Split(bufio.ScanBytes)
		var i int
		for i = 0; s.Scan(); i++ {
			if b := s.Bytes(); len(b) != 1 || b[0] != test[i] {
				t.Errorf("#%d: %d: expected %q got %q", n, i, test, b)
			} else {
				t.Log(string(b))
			}
		}
		if i != len(test) {
			t.Errorf("#%d: termination expected at %d; got %d", n, len(test), i)
		}
		err := s.Err()
		if err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}

var wordScanTests = []string{
	"",
	" ",
	"\n",
	"a",
	" a ",
	"abc def",
	" abc def ",
	" abc\tdef\nghi\rjkl\fmno\vpqr\u0085stu\u00a0\n",
}

// Test that the word splitter returns the same data as strings.Fields.
func TestScanWords(t *testing.T) {
	for n, test := range wordScanTests {
		buf := strings.NewReader(test)
		s := bufio.NewScanner(buf)
		s.Split(bufio.ScanWords)
		words := strings.Fields(test)
		var wordCount int
		for wordCount = 0; wordCount < len(words); wordCount++ {
			if !s.Scan() {
				break
			}
			got := s.Text()
			if got != words[wordCount] {
				t.Errorf("#%d: %d: expected %q got %q", n, wordCount, words[wordCount], got)
			}
		}
		if s.Scan() {
			t.Errorf("#%d: scan ran too long, got %q", n, s.Text())
		}
		if wordCount != len(words) {
			t.Errorf("#%d: termination expected at %d; got %d", n, len(words), wordCount)
		}
		err := s.Err()
		if err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}
