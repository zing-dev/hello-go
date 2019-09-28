package unsafe_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	s := struct {
		a byte
		b byte
		c byte
		d int64 // 8
	}{10, 20, 30, 40}

	s = struct {
		a byte
		b byte
		c byte
		d int64 // 8
	}{0, 0, 20, 10}

	fmt.Println(unsafe.Sizeof(struct {
		a int
		b byte
		c int
	}{0, 0, 0})) //16
	fmt.Println(unsafe.Sizeof(s))   //16
	fmt.Println(unsafe.Sizeof(s.a)) //1
	fmt.Println(unsafe.Sizeof(s.b)) //1
	fmt.Println(unsafe.Sizeof(s.c)) //1
	fmt.Println(unsafe.Sizeof(s.d)) //8
	fmt.Println(unsafe.Pointer(&s)) //address

	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用（偏移量为 0）
	up0 := uintptr(p)
	// 将通用指针转换为 byte 型指针
	pb := (*byte)(p)
	// 给转换后的指针赋值
	*pb = 10
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 2 个字段
	up := up0 + unsafe.Offsetof(s.b)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	// 结构体内容跟着改变
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.b))) = 21
	fmt.Println(s)

	// 偏移到第 3 个字段
	up = up0 + unsafe.Offsetof(s.c)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 30
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.c))) = 31
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 4 个字段
	up = up0 + unsafe.Offsetof(s.d)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 int64 型指针
	pi := (*int64)(p)
	// 给转换后的指针赋值
	*pi = 40
	// 结构体内容跟着改变
	fmt.Println(s)
}
