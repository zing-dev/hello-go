package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

type Language struct {
	Name    string `json:"name"`
	name    string
	Age     int    `json:"age"`
	From    string `json:"from"`
	Popular int    `json:"popular"`
}

var golang = &Language{
	Name:    "Golang",
	Age:     12,
	From:    "google",
	Popular: 6666666,
}

func reflect1() {

	typeOf := reflect.TypeOf(golang)
	fmt.Println(typeOf) //*reflect.Language
	fmt.Println(typeOf.Name())
	fmt.Println(typeOf.Elem().Name())
	fmt.Println(typeOf.Elem().NumField())

	field := typeOf.Elem().Field(0)
	fmt.Println(field.Name)
	fmt.Println(field.Type)
	fmt.Println(field.Anonymous)
	fmt.Println(field.Index)
	fmt.Println(field.Tag)
	fmt.Println(field.PkgPath)

	field = typeOf.Elem().Field(1)
	fmt.Println(field.Name)
	fmt.Println(field.Type)
	fmt.Println(field.Anonymous)
	fmt.Println(field.Index)
	fmt.Println(field.Tag)
	fmt.Println(field.PkgPath)

	name, b := typeOf.Elem().FieldByName("Name")
	fmt.Println(name, b)

}

func reflect2() {
	v := reflect.ValueOf(golang).Elem()
	//{Golang  12 google 6666666}
	fmt.Println(v)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		/*
			{Name  string json:"name" 0 [0] false}
			{name github.com/zhangrxiang/hello-go/std/reflect string  16 [1] false}
			{Age  int json:"age" 32 [2] false}
			{From  string json:"from" 40 [3] false}
			{Popular  int json:"popular" 56 [4] false}
		*/
		//fmt.Println(fieldInfo)

		/*
			json:"name"

			json:"age"
			json:"from"
			json:"popular"
		*/

		//fmt.Println(fieldInfo.Tag.Get("json"))
		fmt.Println(fieldInfo.Tag)
	}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		fmt.Println(fieldInfo.Tag)
	}
}

//将结构体里的成员按照json名字来赋值
func reflect3(ptr interface{}, fields map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]
		if value, ok := fields[name]; ok {
			fmt.Println("类型1：", reflect.ValueOf(value).Type(), "类型2：", v.FieldByName(fieldInfo.Name).Type())
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}

		}
	}

	return
}
