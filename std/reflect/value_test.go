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

func TestValueOfMethods(t *testing.T) {
	str := "hello world"
	value := reflect.ValueOf(&str)
	t.Log("String: ", value.String())
	t.Log("CanSet: ", value.CanSet())
	t.Log("IsValid: ", value.IsValid())
	t.Log("Pointer: ", value.Pointer())
	t.Log("Type: ", value.Type())
	t.Log("Addr: ", value.Addr())
}

func TestStruct(t *testing.T) {
	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	user := &User{Id: 1, Name: "zing"}
	value := reflect.ValueOf(user)
	t.Log(user)
	t.Log(value)
	t.Log(value.IsValid())
	t.Log(value.CanSet())
	t.Log(value.Type())
	t.Log(value.Kind())
	elem := value.Elem()
	t.Log(elem.Kind())
	t.Log(elem.Type())
	elem.Field(0).SetInt(2)
	t.Log(user)
	t.Log(elem.Field(1))
}
