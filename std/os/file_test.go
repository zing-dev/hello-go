package os

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestF2(t *testing.T) {
	f2()
}

func TestF3(t *testing.T) {
	f3()
}

func TestF4(t *testing.T) {
	f4()
}

func TestFileStat(t *testing.T) {
	stat, err := os.Stat("file_test.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stat.ModTime())
	data, ok := stat.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		log.Println(time.Unix(data.CreationTime.Nanoseconds()/1e9, 0))
	}
}
