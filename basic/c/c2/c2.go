package c2

//#include "./hello.h"
import "C"

func c() {
	C.SayHello(C.CString("Hello, World\n"))
	C.SayHello(C.CString("Hello, CGO\n"))
}
