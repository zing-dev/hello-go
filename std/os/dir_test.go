package os

import (
	"fmt"
	"os"
	"testing"
)

func TestReaddir(t *testing.T) {
	file, err := os.Open("./")
	if err != nil {
		t.Fatal(err)
	}
	dirs, err := file.Readdir(-1)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range dirs {
		fmt.Println(f.Name())
	}
}

func TestReaddirnames(t *testing.T) {
	file, err := os.Open("./")
	if err != nil {
		t.Fatal(err)
	}
	dirs, err := file.Readdirnames(2)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range dirs {
		fmt.Println(f)
	}
}
