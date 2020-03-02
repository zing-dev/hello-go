package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type People struct{}
type Student struct{}

type Boy struct {
	Name string
}

func StructToBytes() {
	data, err := json.Marshal(People{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data) //[123 125]
	data, err = json.Marshal(People{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data) //[123 125]

	data, err = json.Marshal(Student{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data) //[123 125]

	data, err = json.Marshal(Boy{
		Name: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data) //[123 34 78 97 109 101 34 58 34 34 125]
	data, err = json.Marshal(Boy{
		Name: "zing",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data) //[123 34 78 97 109 101 34 58 34 122 105 110 103 34 125]
	log.Println(string(data))
	log.Println(fmt.Sprintf("%c%c%c%c", 122, 105, 110, 103))
	log.Println(fmt.Sprintf("%c%c", 123, 125))
	log.Println(fmt.Sprintf("%c%c%c%c%c%c%c%c%c", 34, 78, 97, 109, 101, 34, 58, 34, 34))
}

func BytesToStruct() {
	p := People{}
	err := json.Unmarshal([]byte{123, 125}, &p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)

	b := Boy{}
	err = json.Unmarshal([]byte(`{"Name":"zing"}`), &b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)
}
