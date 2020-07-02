package byte

import (
	"bytes"
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
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	log.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))

}
