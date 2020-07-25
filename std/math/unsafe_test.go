package math

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func TestFloat32bits(t *testing.T) {
	fmt.Println(math.Float32bits(1.0)) //1065353216
	fmt.Println(math.Float32bits(1.1)) //1066192077

	var f = float32(1.0)
	fmt.Println(&f)
	fmt.Println(unsafe.Pointer(&f))
	fmt.Println((*uint8)(unsafe.Pointer(&f)))
	fmt.Println((*uint32)(unsafe.Pointer(&f)))
	fmt.Println((*uint64)(unsafe.Pointer(&f)))
	fmt.Println(*(*uint8)(unsafe.Pointer(&f)))
	fmt.Println(*(*uint32)(unsafe.Pointer(&f)))
	fmt.Println(*(*uint64)(unsafe.Pointer(&f)))
}

func TestFloat32frombits(t *testing.T) {
	fmt.Println(math.Float32frombits(uint32(1065353216)))
	fmt.Println(math.Float32frombits(uint32(1066192077)))
	fmt.Println(math.Float32frombits(math.Float32bits(2.2)))
}

func TestFloat64bits(t *testing.T) {
	fmt.Println(math.Float64bits(1.1))
}

func TestFloat64frombits(t *testing.T) {
	fmt.Println(math.Float64frombits(math.Float64bits(math.MaxFloat64)))
	var data = make([]byte, 64)
	binary.BigEndian.PutUint64(data, math.Float64bits(math.MaxFloat64))
	fmt.Println(data)
}
