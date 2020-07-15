package filepath

import (
	"log"
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	log.Println(filepath.Join("./", "test"))
}

func TestExt(t *testing.T) {
	log.Println(filepath.Ext("./test.txt"))
}

func TestAbs(t *testing.T) {
	log.Println(filepath.Abs("./test.txt"))
	log.Println(filepath.Abs("."))
	log.Println(filepath.Abs("/"))
}

func TestSplit(t *testing.T) {
	log.Println(filepath.Split("."))
}
