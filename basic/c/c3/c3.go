package c3

import "C"
import "log"

func c() {
	SayHello(C.CString("hello world"))
	add(1, 2)
}

//export SayHello
func SayHello(s *C.char) {
	log.Print(C.GoString(s))
}

//export add
func add(a, b int) {
	log.Println(a + b)
}
