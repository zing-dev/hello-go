package cpp

// #cgo LDFLAGS: -lstdc++  -lfoo
// #include "bridge.h"
import "C"

func Run() {
	C.bar()
}
