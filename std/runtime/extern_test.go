package runtime

import (
	"log"
	"runtime"
	"testing"
)

func TestCaller(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Println("=========", i, "=============")
		caller, file, line, ok := runtime.Caller(i)
		log.Println(caller, file, line, ok)
	}
}
