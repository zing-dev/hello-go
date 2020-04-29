package basic

import (
	"encoding/json"
	"fmt"
	"log"
)

type Map struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func map1() {
	m := make(map[string]*Map)
	m["1"] = &Map{
		Id:   1,
		Name: "zing",
	}
	m["2"] = &Map{
		Id:   2,
		Name: "zing2",
	}

	for _, v := range m {
		v.Name += "-zing"
	}

	for _, v := range m {
		log.Println(v.Name)
	}

	m2 := make(map[string]Map)
	m2["1"] = Map{
		Id:   1,
		Name: "zing",
	}
	m2["2"] = Map{
		Id:   2,
		Name: "zing2",
	}

	for _, v := range m2 {
		v.Name += "-zing"
	}

	for _, v := range m2 {
		log.Println(v.Name)
	}
}

func map2() {
	m := Map{
		Id:   1,
		Name: "zing",
	}
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	//{"id":1,"name":"zing"}
	fmt.Println(string(data))
	m2 := make(map[string]Map)
	m2["1"] = m
	m.Id = 2
	m2["2"] = m
	data, err = json.Marshal(m2)
	if err != nil {
		log.Fatal(err)
	}
	//{"1":{"id":1,"name":"zing"},"2":{"id":2,"name":"zing"}}
	fmt.Println(string(data))

	m3 := make(map[string]*Map)
	m3["1"] = &Map{
		Id:   1,
		Name: "zing",
	}
	m3["2"] = &Map{
		Id:   2,
		Name: "zing2",
	}
	data, err = json.Marshal(m3)
	if err != nil {
		log.Fatal(err)
	}
	//{"1":{"id":1,"name":"zing"},"2":{"id":2,"name":"zing2"}}
	fmt.Println(string(data))
}

func map3() {
	data, err := json.MarshalIndent(map[string]interface{}{
		"root": map[string]interface{}{
			"students": map[string]interface{}{
				"1": map[string]interface{}{
					"name": "zing",
					"age":  20,
				},
				"2": map[string]interface{}{
					"name": "zing",
					"age":  23,
				},
			},
		},
	}, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func map4() {
	m := make(map[interface{}]interface{})
	m[0] = "0"
	m[0.0] = 0.1
	m["1"] = 1
	m[struct{}{}] = "struct"
	m[false] = false
	m[true] = true

	//map[false:false true:true 0:0.1 0:0 1:1]
	fmt.Println(m)
	fmt.Println(m[0])
	fmt.Println(m[0.0])
	fmt.Println(m["1"])
	fmt.Println(m[struct{}{}])
}
