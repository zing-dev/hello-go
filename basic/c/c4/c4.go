package c4

// extern void SayHello(char* s);
import "C"
import "fmt"

func c() {
	C.SayHello(C.CString("Hello, World\n"))
}

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
