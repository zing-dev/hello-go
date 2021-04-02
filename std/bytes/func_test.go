package bytes

import (
	"bytes"
	"fmt"
	"log"
	"testing"
	"unicode"
)

func TestEqual(t *testing.T) {
	log.Println(bytes.Equal([]byte{1, 2}, []byte{1, 2}))
	log.Println(bytes.Equal([]byte(""), []byte("")))
}

func TestCompare(t *testing.T) {
	log.Println(bytes.Compare([]byte{}, []byte("")))
	log.Println(bytes.Compare([]byte{1}, []byte("1")))
}

func TestCount(t *testing.T) {
	log.Println(bytes.Count([]byte{1, 2, 3}, []byte{1}))
	log.Println(bytes.Count([]byte("cheese"), []byte("e")))
	// If sep is an empty slice, Count returns 1 + the number
	// of UTF-8-encoded code points in s.
	log.Println(bytes.Count([]byte("five"), []byte("")))
}

func TestContains(t *testing.T) {
	log.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	log.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	log.Println(bytes.Contains([]byte("seafood"), []byte("")))
	log.Println(bytes.Contains([]byte(""), []byte("")))
}

func TestContainsAny(t *testing.T) {
	log.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))   //true
	log.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的.")) //true
	log.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	log.Println(bytes.ContainsAny([]byte(""), ""))
}

func TestContainsRune(t *testing.T) {
	log.Println(bytes.ContainsRune([]byte("I like seafood."), 'f'))
	log.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	log.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	log.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!'))
	log.Println(bytes.ContainsRune([]byte(""), '@'))
}

func TestIndexByte(t *testing.T) {
	log.Println(bytes.IndexByte([]byte("chicken"), byte('k')))
	log.Println(bytes.IndexByte([]byte("chicken"), byte('g')))
}

func TestLastIndex(t *testing.T) {
	log.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	log.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	log.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}

func TestLastIndexByte(t *testing.T) {
	log.Println(bytes.LastIndexByte([]byte("go gopher"), byte('g')))
	log.Println(bytes.LastIndexByte([]byte("go gopher"), byte('r')))
	log.Println(bytes.LastIndexByte([]byte("go gopher"), byte('z')))
}

func TestIndexRune(t *testing.T) {
	log.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	log.Println(bytes.IndexRune([]byte("chicken"), 'd'))
}

func TestIndexAny(t *testing.T) {
	log.Println(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	log.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))
}

func TestLastIndexAny(t *testing.T) {
	log.Println(bytes.LastIndexAny([]byte("chicken"), "aeiouy"))
	log.Println(bytes.LastIndexAny([]byte("crwth"), "aeiouy"))
}

func TestSplitN(t *testing.T) {
	log.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	log.Printf("%q (nil = %v)\n", z, z == nil)
}

func TestSplitAfterN(t *testing.T) {
	log.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
}

func TestSplit(t *testing.T) {
	log.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	log.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	log.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	log.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}

func TestSplitAfter(t *testing.T) {
	log.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
}

func TestFields(t *testing.T) {
	log.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}

func TestFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		log.Println(c, !unicode.IsLetter(c) && !unicode.IsNumber(c))
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	log.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
}

func TestJoin(t *testing.T) {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	log.Printf("%s\n", bytes.Join(s, []byte(", ")))
}

func TestHasPrefix(t *testing.T) {
	log.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	log.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	log.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))
	log.Println(string([]byte("")[0:0]) == "")
}

func TestHasSuffix(t *testing.T) {
	log.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	log.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	log.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	log.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}

func TestMap(t *testing.T) {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	log.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))

	log.Println(bytes.Map(func(r rune) rune {
		return r - 'a'
	}, []byte("asd")))
}

func TestRepeat(t *testing.T) {
	log.Println(bytes.Repeat([]byte("a"), 10))
}

func TestToUpper(t *testing.T) {
	log.Println(string(bytes.ToUpper([]byte("asd"))))
	log.Println(string(bytes.ToUpper([]byte("α"))))
	log.Println(string(bytes.ToUpper([]byte("あ"))))
}

func TestToTitle(t *testing.T) {
	log.Println(string(bytes.ToTitle([]byte("hello world"))))
}

func TestTitle(t *testing.T) {
	log.Println(string(bytes.ToTitle([]byte("hello world"))))
}

func TestTrimLeftFunc(t *testing.T) {
	log.Println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter)))
	log.Println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct)))
	log.Println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

func TestTrimRightFunc(t *testing.T) {
	log.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	log.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	log.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

func TestTrimFunc(t *testing.T) {
	log.Println(string(bytes.TrimFunc([]byte("go-gopher"), unicode.IsLetter)))
	log.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	log.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

func TestTrimPrefix(t *testing.T) {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	log.Printf("Hello%s", b)
}

func TestTrimSuffix(t *testing.T) {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	log.Println(b)
}

func TestIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	log.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	log.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}

func TestLastIndexFunc(t *testing.T) {
	log.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	log.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	log.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))
}

func TestTrim(t *testing.T) {
	log.Println(string(bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! ")))
	log.Println(string(bytes.Trim([]byte("!as!!df!!"), "!")))
	log.Println(string(bytes.Trim([]byte("11!as!!df!1!1"), "1!")))
}

func TestTrimLeft(t *testing.T) {
	log.Print(string(bytes.TrimLeft([]byte("453gopher8257"), "0123456789")))
}

func TestTrimRight(t *testing.T) {
	log.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))
}

func TestTrimSpace(t *testing.T) {
	log.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
}

func TestRunes(t *testing.T) {
	rs := bytes.Runes([]byte("go gopher 你好"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}
	log.Printf("%c\n", 0x597D)
}

func TestReplace(t *testing.T) {
	log.Printf("%s\n",
		bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	log.Printf("%s\n",
		bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}

func TestReplaceAll(t *testing.T) {
	log.Printf("%s\n",
		bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
}

func TestEqualFold(t *testing.T) {
	log.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}

func TestIndex(t *testing.T) {
	log.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	log.Println(bytes.Index([]byte("chicken"), []byte("dmr")))
}
