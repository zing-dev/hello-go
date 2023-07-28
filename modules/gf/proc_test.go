package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"log"
	"testing"
)

func TestProc(t *testing.T) {
	result, err := gproc.ShellExec(gctx.New(), "tasklist  /FI 'IMAGENAME eq msedge.exe' /FO csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
