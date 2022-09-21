package reflect_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type I interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type Coordinate struct {
	long float32
	lat  float32
}

type User struct {
	ID         int
	Name       string
	age        int
	address    []string
	Coordinate Coordinate
}

func (u *User) SetAge(age int) *User {
	u.age = age
	return u
}

func TestTypeOf(t *testing.T) {
	fmt.Println(reflect.TypeOf("str"))      //string
	fmt.Println(reflect.TypeOf(0))          //int
	fmt.Println(reflect.TypeOf(int32(0)))   //int32
	fmt.Println(reflect.TypeOf(0.0))        //float64
	fmt.Println(reflect.TypeOf(true))       //bool
	fmt.Println(reflect.TypeOf(true))       //bool
	fmt.Println(reflect.TypeOf(nil))        //bool
	fmt.Println(reflect.TypeOf(nil) == nil) //bool
	fmt.Println(reflect.TypeOf(struct{}{})) //struct {}

	tStr := reflect.TypeOf("")
	fmt.Println(tStr.Name())
	fmt.Println(tStr.Size())

	tArr := reflect.TypeOf([...]int{1, 2, 3, 4, 5})
	fmt.Println(tArr.Name())        //""
	fmt.Println(tArr.Size())        //40
	fmt.Println(tArr.Kind())        //array
	fmt.Println(tArr.Len())         //5
	fmt.Println(tArr.Align())       //8
	fmt.Println(tArr.Elem())        //int
	fmt.Println(tArr.Elem().Kind()) //int

	tFun := reflect.TypeOf(func(a, b int) int {
		return 0
	})
	fmt.Println(tFun.Name())   //""
	fmt.Println(tFun.Kind())   //func
	fmt.Println(tFun.Align())  //8
	fmt.Println(tFun.NumIn())  //2
	fmt.Println(tFun.NumOut()) //1
	fmt.Println(tFun.In(0))    //int
	fmt.Println(tFun.In(1))    //int
	fmt.Println(tFun.Out(0))   //int

	fmt.Println(tFun.NumMethod())  //0
	fmt.Println(tFun.IsVariadic()) //false

	tS := reflect.TypeOf(User{
		Name: "zing",
		age:  25,
	})
	fmt.Println(tS.Name()) //User
	fmt.Println(tS.Kind()) //struct
	fmt.Println(tS.Align())
	fmt.Println(tS.NumField())
	fmt.Println(tS.NumMethod())
	fmt.Println(tS.PkgPath())
	fmt.Println(tS.Field(0))
	fmt.Println(tS.Field(1))
	fmt.Println(tS.FieldByName("age"))
}

func TestKind(t *testing.T) {
	fmt.Println(reflect.Kind(1))
	fmt.Println(reflect.Kind(0.0))
}

type S struct {
	A    string `json:"a"`
	B    byte   `json:"b"`
	C    bool   `json:"c,omitempty"`
	Name string `json:"name"`
}

func (s S) M1() string {
	return s.A
}

func TestTag(t *testing.T) {
	ts := reflect.TypeOf(*&S{
		A:    "hello world",
		B:    1,
		C:    false,
		Name: "zing",
	})
	fmt.Println(ts)

	fmt.Println(ts.NumField())
	fmt.Println(ts.NumMethod())
	for i := 0; i < ts.NumField(); i++ {
		if value, ok := ts.Field(i).Tag.Lookup("json"); ok {
			fmt.Println(value, strings.Split(value, ","), strings.Split(value, ",")[0])
		}
		fmt.Println("Field.Name: ", ts.Field(i).Name)
		fmt.Println("Field.Type: ", ts.Field(i).Type)
	}

}
