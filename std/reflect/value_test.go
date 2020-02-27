package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	fmt.Println(reflect.ValueOf("hello world"))
	fmt.Println(reflect.ValueOf(true))
	fmt.Println(reflect.ValueOf(1))
	fmt.Println(reflect.ValueOf(1.1))
	fmt.Println(reflect.ValueOf([]byte{}))
	fmt.Println(reflect.ValueOf([...]byte{1, 2, 3}))
	fmt.Println(reflect.ValueOf(map[string]string{}))
	fmt.Println(reflect.ValueOf(make(chan int)))

	fmt.Println(reflect.ValueOf("hello").String())
	fmt.Println(reflect.ValueOf(true).Bool())
	fmt.Println(reflect.ValueOf(1).Int())
}
