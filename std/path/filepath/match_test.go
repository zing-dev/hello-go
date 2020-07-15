package filepath

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMatch(t *testing.T) {
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
}

func TestGlob(t *testing.T) {
	fmt.Println(filepath.Glob("./"))
	fmt.Println(filepath.Glob("/"))
	fmt.Println(filepath.Glob("C:"))
}
