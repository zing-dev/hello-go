package parser

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"
)

func TestStart(t *testing.T) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "parser_test.go", nil, parser.ImportsOnly)
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range file.Imports {
		fmt.Println(i.Path.Value)
	}
}
