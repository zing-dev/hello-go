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