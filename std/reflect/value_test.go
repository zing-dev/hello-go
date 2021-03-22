package reflect_test

import (
	"fmt"
	"reflect"
	"runtime"
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

func TestIndirect(t *testing.T) {
	s := "hello"
	t.Log(reflect.Indirect(reflect.ValueOf(s)))
	t.Log(reflect.Indirect(reflect.ValueOf(s)).String() == s)
	s2 := &s
	t.Log(reflect.Indirect(reflect.ValueOf(s2)))
	t.Log(reflect.Indirect(reflect.ValueOf(s2)).String() == s)
	var s3 *string
	t.Log(reflect.Indirect(reflect.ValueOf(s3))) //<invalid reflect.Value>
}

func TestMakeSlice(t *testing.T) {
	var i []int
	value := reflect.MakeSlice(reflect.TypeOf(i), 10, 10)
	t.Log(value.String())

	value.Index(0).SetInt(10)
	value.Index(1).SetInt(20)
	t.Log(value.String())
	t.Log(value.Index(0))
	t.Log(value.Index(1))
}

func TestSliceOf(t *testing.T) {
	// check construction and use of type not in binary
	type T int
	st := reflect.SliceOf(reflect.TypeOf(T(1)))
	if got, want := st.String(), "[]reflect_test.T"; got != want {
		t.Errorf("SliceOf(T(1)).String()=%q, want %q", got, want)
	}
	v := reflect.MakeSlice(st, 10, 10)
	runtime.GC()
	for i := 0; i < v.Len(); i++ {
		v.Index(i).Set(reflect.ValueOf(T(i)))
		runtime.GC()
	}
	s := fmt.Sprint(v.Interface())
	want := "[0 1 2 3 4 5 6 7 8 9]"
	if s != want {
		t.Errorf("constructed slice = %s, want %s", s, want)
	}

	// check that type already in binary is found
	type T1 int
	checkSameType(t, reflect.SliceOf(reflect.TypeOf(T1(1))), []T1{})
}

func checkSameType(t *testing.T, x reflect.Type, y interface{}) {
	if x != reflect.TypeOf(y) || reflect.TypeOf(reflect.Zero(x).Interface()) != reflect.TypeOf(y) {
		t.Errorf("did not find preexisting type for %s (vs %s)", reflect.TypeOf(x), reflect.TypeOf(y))
	}
}
