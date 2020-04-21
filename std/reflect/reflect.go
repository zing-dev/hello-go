package reflect

import (
	"fmt"
	"reflect"
)

type Language struct {
	Name    string `json:"name"`
	name    string
	Age     int `json:"age"`
	From    string
	Popular int
}

func reflect1() {
	golang := &Language{
		Name:    "Golang",
		Age:     12,
		From:    "google",
		Popular: 6666666,
	}
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
