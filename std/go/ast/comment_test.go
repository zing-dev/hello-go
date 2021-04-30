package ast

import (
	"bufio"
	"go/ast"
	"io"
	"log"
	"os"
	"testing"
)

// 这是一个测试函数,用来解析注释,获取注释文本
// this is comment
func TestCommentStart(t *testing.T) {
	file, err := os.Open("comment_test.go")
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(file)
	g := make([]*ast.Comment, 0)
	for reader.Size() > 0 {
		data, err := reader.ReadBytes('\n')
		if len(data) != 0 {
			g = append(g, &ast.Comment{Text: string(data)})
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	group := ast.CommentGroup{List: g}
	log.Println(group.Text())
}
