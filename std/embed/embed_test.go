package main_test

import (
	"embed"
	"fmt"
	"io"
	"testing"
)

//go:embed *.go *.mod *tmpl
//go:embed www
var content embed.FS

func TestName(t *testing.T) {
	dir, _ := content.ReadDir("www")

	for _, entry := range dir {
		entry.Name()
	}
	file, err := content.Open("www/1.txt")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}
