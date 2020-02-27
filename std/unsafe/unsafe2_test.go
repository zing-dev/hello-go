package unsafe_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeof(t *testing.T) {
	fmt.Println(unsafe.Sizeof(true))                           //1
	fmt.Println(unsafe.Sizeof(int8(0)))                        //1
	fmt.Println(unsafe.Sizeof(int16(0)))                       //2
	fmt.Println(unsafe.Sizeof(int32(0)))                       //4
	fmt.Println(unsafe.Sizeof(int64(0)))                       //8
	fmt.Println(unsafe.Sizeof(0))                              //8
	fmt.Println(unsafe.Sizeof(int(0)))                         //8
	fmt.Println(unsafe.Sizeof(float32(0)))                     //4
	fmt.Println(unsafe.Sizeof(float64(0)))                     //8
	fmt.Println(unsafe.Sizeof(0.0))                            //8
	fmt.Println(unsafe.Sizeof(""))                             //16
	fmt.Println(unsafe.Sizeof("hello world"))                  //16
	fmt.Println(unsafe.Sizeof([]byte{}))                       //24
	fmt.Println(unsafe.Sizeof([]byte{1}))                      //24
	fmt.Println(unsafe.Sizeof([]byte{1, 2}))                   //24
	fmt.Println(unsafe.Sizeof([]byte{1, 2, 3}))                //24
	fmt.Println(unsafe.Sizeof([...]byte{1, 2, 3}))             //3
	fmt.Println(unsafe.Sizeof([3]byte{1}))                     //3
	fmt.Println(unsafe.Sizeof([...]int{1, 2, 3}))              //24
	fmt.Println(unsafe.Sizeof([3]int{}))                       //24
	fmt.Println(unsafe.Sizeof([]int8{}))                       //24
	fmt.Println(unsafe.Sizeof(map[string]int{}))               //8
	fmt.Println(unsafe.Sizeof(map[string]int{"1": 1, "2": 2})) //8
	fmt.Println(unsafe.Sizeof(struct {
	}{})) //0
	fmt.Println(unsafe.Sizeof(struct {
		a int
	}{})) //8
	fmt.Println(unsafe.Sizeof(struct {
		a int
		b byte
	}{})) //16

	fmt.Println(unsafe.Sizeof(struct {
		i8  int8  //1
		i16 int16 //2
		i32 int32 //4
	}{})) //8

	fmt.Println(unsafe.Sizeof(struct {
		i8  int8  //1
		i32 int32 //4
		i16 int16 //2
		i64 int64 //8
	}{})) //24 = 4 + 4 + 8 + 8
	fmt.Println(unsafe.Sizeof(struct {
		i8  int8  //1
		i16 int16 //2
		i32 int32 //4
		i64 int64 //8
	}{})) //16 = 2+2+4+8

	fmt.Println(unsafe.Sizeof(struct {
		i64 int64 //8
		i16 int16 //2
		i32 int32 //4
		i8  int8  //1
	}{})) //24 = 8 + 4 + 4 + 8
}
