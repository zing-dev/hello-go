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
		{K: "ok1", V: true},
		{K: "ok2", V: Student{}},
		{K: "ok3", V: Route{
			Name:   "api",
			Method: "GET",
			Ok:     true,
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
			route, ok := a.V.(*Route)
			log.Println(ok, route)
			r, ok := a.V.(map[string]any)
			log.Println(ok, r)
			s, ok := r["name"].(string)
			log.Println(ok, s)
			o, ok := r["ok"].(bool)
			log.Println(ok, o)

			data, err := json.Marshal(a.V)
			if err != nil {
				log.Fatal(err)
			}
			route2 := new(Route)
			err = json.Unmarshal(data, route2)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(*route2)
		}
	}
}

func MapToStruct(m, s any) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}
	return err
}
