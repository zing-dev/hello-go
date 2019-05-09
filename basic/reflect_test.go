package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type Human struct {
	username string
	age      int
}

func TestReflectTypeOf(t *testing.T) {

	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(1.0))
	fmt.Println(reflect.TypeOf('a'))
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf([]byte{}))
	fmt.Println(reflect.TypeOf([...]byte{0, 1}))
	fmt.Println(reflect.TypeOf([...]string{"aaa", "bbb"}))
	var a interface{}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(map[string]string{"aaa": "bbb"}))
	fmt.Println(reflect.TypeOf(map[string]map[string]string{}))

	fmt.Println(reflect.TypeOf([]string{}))
	fmt.Println(reflect.TypeOf([]map[string]string{}))
	fmt.Println(reflect.TypeOf([...]map[string]string{{"aaa": "bbb"}}))
	fmt.Println(reflect.TypeOf([]map[string]map[string]string{}))

	fmt.Println(reflect.TypeOf(Human{"zing", 25}))

}

func TestReflectValueOf(t *testing.T) {
	fmt.Println(reflect.ValueOf(100))
	fmt.Println(reflect.ValueOf(1.0))
	fmt.Println(reflect.ValueOf('a'))

	fmt.Println(reflect.ValueOf([...]byte{0, 1}))
	fmt.Println(reflect.ValueOf(map[string]string{"aaa": "bbb"}))

	fmt.Println(reflect.ValueOf(Human{"zing", 25}))

}
