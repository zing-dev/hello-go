// +build go1.10

package c5

// extern void SayHello(_GoString_ s);
import "C"
import "fmt"

func c() {
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
