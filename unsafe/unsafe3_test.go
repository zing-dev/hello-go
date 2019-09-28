package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	age  int
	name string
}

func TestUnsafeStruct(t *testing.T) {

	user := &User{}
	fmt.Println(user)

	s := (*int)(unsafe.Pointer(user))
	*s = 10

	up := uintptr(unsafe.Pointer(user)) + unsafe.Sizeof(int(0))

	namep := (*string)(unsafe.Pointer(up))
	*namep = "zing"

	fmt.Println(user)

	u := unsafe.Alignof(user) //8
	fmt.Println(u)
}

func TestAlignof(t *testing.T) {
	type S struct {
		a int     //8
		b byte    //1
		c string  //16
		d float64 //8
	}
	fmt.Println(unsafe.Alignof(S{})) //8
	fmt.Println(unsafe.Sizeof(S{}))  //40 = 8 + 8 +16 +8
}
