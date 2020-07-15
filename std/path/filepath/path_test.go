package filepath

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestClean(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(dir)
	}
	log.Println(dir)
	log.Println(filepath.Clean(dir))
	log.Println(filepath.Clean(""))                 //.
	log.Println(filepath.Clean("."))                //.
	log.Println(filepath.Clean("/"))                // C:\
	log.Println(filepath.Clean("C://"))             //\
	log.Println(filepath.Clean("./test.txt"))       // test.txt
	log.Println(filepath.Clean("./../../test.txt")) //..\..\test.txt
}

func TestJoin(t *testing.T) {
	log.Println(filepath.Join("./", "test"))          //test
	log.Println(filepath.Join("./../", "./test.txt")) //..\test.txt
}

func TestExt(t *testing.T) {
	log.Println(filepath.Ext("./test.txt")) //.txt
	log.Println(filepath.Ext("./test"))     //
}

func TestAbs(t *testing.T) {
	log.Println(filepath.Abs("./test.txt"))
	log.Println(filepath.Abs("."))
	log.Println(filepath.Abs("/"))
}

func TestSplit(t *testing.T) {
	log.Println(filepath.Split("."))
	log.Println(filepath.Split("../../../test.txt"))

	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}
	log.Println("On Unix:")
	for _, p := range paths {
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
}

func TestToSlash(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(dir)
	}
	log.Println(dir)                   //C:\Users\zhang\workspace\go\hello\hello-go\std\path\filepath
	log.Println(filepath.ToSlash(dir)) //C:/Users/zhang/workspace/go/hello/hello-go/std/path/filepath
}

func TestFromSlash(t *testing.T) {
	log.Println(filepath.FromSlash("C:/Users/zhang/workspace/go/hello/hello-go/std/path/filepath"))
}

func TestSplitList(t *testing.T) {
	log.Println(filepath.SplitList("./../../test.txt"))
	log.Println(filepath.SplitList(".\\..\\..\\test.txt"))
	log.Println(fmt.Sprintf("%c", filepath.ListSeparator)) //;
	log.Println(fmt.Sprintf("%c", filepath.Separator))     //\
	list := filepath.SplitList("../../;./;/")
	log.Println(len(list))
	log.Println(list)
}

func TestEvalSymlinks(t *testing.T) {
	log.Println(filepath.EvalSymlinks("./test.txt"))
	log.Println(filepath.EvalSymlinks("./filepath_test.go"))
}

func TestRel(t *testing.T) {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}
}

func TestWalk(t *testing.T) {
	err := filepath.Walk("./../", func(path string, info os.FileInfo, err error) error {
		log.Println(path, info.Name())
		return nil
	})
	log.Println(err)
}

func TestBase(t *testing.T) {
	log.Println(filepath.Base("./././././"))
	log.Println(filepath.Base("../../../"))
	log.Println(filepath.Base("./../test.txt"))
}

func TestDir(t *testing.T) {
	log.Println(filepath.Dir("./././././"))
	log.Println(filepath.Dir("../../../"))
	log.Println(filepath.Dir("./../test.txt"))
}

func TestVolumeName(t *testing.T) {
	log.Println(filepath.VolumeName("./../test.txt"))
	log.Println(filepath.VolumeName("C:\\foo\\bar"))
	log.Println(filepath.VolumeName("\\host\\share\\foo"))
}
