package runtime

import (
	"log"
	"runtime"
	"testing"
)

func Caller(skip int) (uintptr, string, int, bool) {
	return runtime.Caller(skip)
}

func TestCaller(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Println("=========", i, "=============")
		caller, file, line, ok := Caller(i)
		log.Println(caller, file, line, ok)
	}
}
