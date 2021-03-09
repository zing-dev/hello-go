package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestReflectString(t *testing.T) {
	s := "hello world"
	s1 := s + "hello world"
	s2 := s1 + "hello world"
	s3 := (*reflect.StringHeader)(unsafe.Pointer(&s))
	t.Log(*(*string)(unsafe.Pointer(s3)))
	t.Log(*&s)
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
}
