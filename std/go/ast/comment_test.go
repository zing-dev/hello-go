package ast

import (
	"bufio"
	"fmt"
	"go/ast"
	"io"
	"os"
	"testing"
)

func TestCommentLine(t *testing.T) {
	group := ast.CommentGroup{List: []*ast.Comment{
		{Text: "// 这是一个测试函数,用来解析注释,获取注释文本"},
		{Text: "// hello"},
		{Text: "//hello"},
		{Text: "/*hello*/"},
		{Text: "/* hello\n\r\thello */ \n// nihao"},
	}}
	fmt.Println(group.Text())
}

// 这是一个测试函数,用来解析注释,获取注释文本
// this is comment
func TestCommentFile(t *testing.T) {
	data, err := os.ReadFile("comment_test.go")
	if err != nil {
		t.Fatal(err)
	}
	group := &ast.CommentGroup{List: []*ast.Comment{{Text: string(data)}}}
	fmt.Println(group.Text())

}

func TestCommentStart(t *testing.T) {
	file, err := os.Open("comment_test.go")
	if err != nil {
		t.Fatal(err)
	}
	reader := bufio.NewReader(file)
	list := make([]*ast.Comment, 0)
	for reader.Size() > 0 {
		line, _, err := reader.ReadLine()
		text := string(line)
		list = append(list, &ast.Comment{Text: text})
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	group := ast.CommentGroup{List: list}
	fmt.Println(group.Text())
}
