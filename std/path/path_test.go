package path

import (
	"fmt"
	"log"
	"path"
	"testing"
)

func TestClean(t *testing.T) {
	log.Println(path.Clean("."))          //.
	log.Println(path.Clean("./"))         //.
	log.Println(path.Clean(".."))         //..
	log.Println(path.Clean("../"))        //..
	log.Println(path.Clean("./../"))      //..
	log.Println(path.Clean("../../"))     //../..
	log.Println(path.Clean("std/path/"))  //std/path
	log.Println(path.Clean("/std/path/")) //std/path
	log.Println(path.Clean("std\\path\\"))
	log.Println(path.Clean("std\\path\\path_test.go"))
	log.Println(path.Clean("std/path/path_test.go"))
	log.Println(path.Clean("C:\\Users\\zing\\workspace\\go\\hello\\hello-go\\std\\path\\path_test.go"))
}

func TestSplit(t *testing.T) {
	log.Println(path.Split("C:\\Users\\zing\\workspace\\go\\hello\\hello-go\\std\\path\\path_test.go"))
	log.Println(path.Split("std/path/path_test.go"))
	log.Println(path.Split("c:/std/path/path_test.go"))
	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css")
	split("myfile.css")
	split("")
}

func TestJoin(t *testing.T) {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
	fmt.Println(path.Join("/a/b", "c.txt"))
	fmt.Println(path.Join("/a/b/", "c.txt"))
	fmt.Println(path.Join("/a/b/", "/c.txt"))
	fmt.Println(path.Join("../../a/b/", "../c.txt")) //../../a/c.txt
}

func TestExt(t *testing.T) {
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))
}

func TestBase(t *testing.T) {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

func TestIsAbs(t *testing.T) {
	fmt.Println(path.IsAbs("/dev/null"))
}

func TestDir(t *testing.T) {
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
}
