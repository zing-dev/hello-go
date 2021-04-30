package ast

import (
	"go/ast"
	"strings"
	"testing"
)

var comments = []struct {
	list []string
	text string
}{
	{[]string{"//"}, ""},
	{[]string{"//   "}, ""},
	{[]string{"//", "//", "//   "}, ""},
	{[]string{"// foo   "}, "foo\n"},
	{[]string{"//", "//", "// foo"}, "foo\n"},
	{[]string{"// foo  bar  "}, "foo  bar\n"},
	{[]string{"// foo", "// bar"}, "foo\nbar\n"},
	{[]string{"// foo", "//", "//", "//", "// bar"}, "foo\n\nbar\n"},
	{[]string{"// foo", "/* bar */"}, "foo\n bar\n"},
	{[]string{"//", "//", "//", "// foo", "//", "//", "//"}, "foo\n"},

	{[]string{"/**/"}, ""},
	{[]string{"/*   */"}, ""},
	{[]string{"/**/", "/**/", "/*   */"}, ""},
	{[]string{"/* Foo   */"}, " Foo\n"},
	{[]string{"/* Foo  Bar  */"}, " Foo  Bar\n"},
	{[]string{"/* Foo*/", "/* Bar*/"}, " Foo\n Bar\n"},
	{[]string{"/* Foo*/", "/**/", "/**/", "/**/", "// Bar"}, " Foo\n\nBar\n"},
	{[]string{"/* Foo*/", "/*\n*/", "//", "/*\n*/", "// Bar"}, " Foo\n\nBar\n"},
	{[]string{"/* Foo*/", "// Bar"}, " Foo\nBar\n"},
	{[]string{"/* Foo\n Bar*/"}, " Foo\n Bar\n"},

	{[]string{"// foo", "//go:noinline", "// bar", "//:baz"}, "foo\nbar\n:baz\n"},
	{[]string{"// foo", "//lint123:ignore", "// bar"}, "foo\nbar\n"},
}

func TestCommentText(t *testing.T) {
	for i, c := range comments {
		list := make([]*ast.Comment, len(c.list))
		for i, s := range c.list {
			list[i] = &ast.Comment{Text: s}
		}

		text := (&ast.CommentGroup{List: list}).Text()
		if text != c.text {
			t.Errorf("case %d: got %q; expected %q", i, text, c.text)
		} else {
			t.Logf("case %d: got %q; expected %q", i, text, c.text)
		}
	}
}

var isDirectiveTests = []struct {
	in string
	ok bool
}{
	{"abc", false},
	{"go:inline", true},
	{"Go:inline", false},
	{"go:Inline", false},
	{":inline", false},
	{"lint:ignore", true},
	{"lint:1234", true},
	{"1234:lint", true},
	{"go: inline", false},
	{"go:", false},
	{"go:*", false},
	{"go:x*", true},
}

func TestIsDirective(t *testing.T) {
	for _, tt := range isDirectiveTests {
		if ok := isDirective(tt.in); ok != tt.ok {
			t.Errorf("isDirective(%q) = %v, want %v", tt.in, ok, tt.ok)
		}
	}
}

// isDirective reports whether c is a comment directive.
func isDirective(c string) bool {
	// "//line " is a line directive.
	// (The // has been removed.)
	if strings.HasPrefix(c, "line ") {
		return true
	}

	// "//[a-z0-9]+:[a-z0-9]"
	// (The // has been removed.)
	colon := strings.Index(c, ":")
	if colon <= 0 || colon+1 >= len(c) {
		return false
	}
	for i := 0; i <= colon+1; i++ {
		if i == colon {
			continue
		}
		b := c[i]
		if !('a' <= b && b <= 'z' || '0' <= b && b <= '9') {
			return false
		}
	}
	return true
}
