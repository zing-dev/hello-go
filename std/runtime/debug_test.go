package runtime

import (
	"runtime"
	"testing"
)

func TestGOMAXPROCS(t *testing.T) {
	t.Log(runtime.GOMAXPROCS(0))
	t.Log(runtime.GOMAXPROCS(2))
	t.Log(runtime.GOMAXPROCS(1))
}

func TestNumCPU(t *testing.T) {
	t.Log(runtime.NumCPU())
}

func TestNumCgoCall(t *testing.T) {
	t.Log(runtime.NumCgoCall())
}

func TestNumGoroutine(t *testing.T) {
	go func() {}()
	t.Log(runtime.NumGoroutine())
	runtime.LockOSThread()
}
