package json

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Attr struct {
	K string `json:"k"`
	V any    `json:"v"`
}

func TestUnmarshal() {
	for _, attr := range []Attr{
		{K: "name", V: "zing"},
		{K: "age", V: 22},
		{K: "ok", V: true},
		{K: "ok", V: Student{}},
		{K: "ok", V: Route{
			Name:   "api",
			Method: "GET",
		}},
	} {
		data, err := json.Marshal(attr)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(data))
		var a = new(Attr)
		err = json.Unmarshal(data, a)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("%#v - %#v  %s", a.K, a.V, reflect.TypeOf(a.V)))
		typeOf := reflect.TypeOf(a.V)
		if typeOf.Kind() == reflect.Map {
			log.Println("map")
		}

	}

}
