package runtime

import (
	"log"
	"runtime"
	"testing"
)

func TestFuncForPC(t *testing.T) {
	pc, file, line, ok := runtime.Caller(0)
	log.Println(pc, file, line, ok)

	forPC := runtime.FuncForPC(pc)
	log.Println(forPC.Name()) //github.com/zing-dev/hello-go/std/runtime.TestFuncForPC
	fileLine, l := forPC.FileLine(pc)
	log.Println(fileLine, l)
}
