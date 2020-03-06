package c6

import "C"
import (
	"log"
	"unsafe"
)

//https://chai2010.cn/gopherchina2018-cgo-talk/#/6/8

//int32 和 *C.char 相互转换
//第一步: int32 => uintprt
//第二步: uintptr => unsafe.Pointer
//第三步: unsafe.Pointer => *C.char
func IToC() {
	// int32 => *C.char
	var x = int32(9527)
	var p *C.char = (*C.char)(unsafe.Pointer(uintptr(x)))
	//log.Println(C.GoString(p))
	_ = p
	// *C.char => int32
	var y *C.char
	var q int32 = int32(uintptr(unsafe.Pointer(y)))
	log.Println(q)
}
