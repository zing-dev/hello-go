package json

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type People struct{}
type Student struct{}

type Boy struct {
	Name string
}
type Version struct {
	Version   string `json:"version"`
	Log       string `json:"log"`
	Status    string `json:"status"`
	GitHash   string `json:"git_hash"`
	updatedAt string
	hash      string
}

type Soft struct {
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	Author    string    `json:"author"`
	Version   []Version `json:"versions"`
	Copyright string    `json:"copyright"`
	Inherit   bool      `json:"inherit"`
}

type Route struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Ok     bool   `json:"ok"`
}

var R = map[Route]string{
	Route{
		Name:   "system/sms",
		Method: "post",
	}: "保存短信配置",
	Route{
		Name:   "system/sms",
		Method: "post",
	}: "保存短信配置",
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

func JSON() {
	type Stu struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data, _ := json.Marshal(Stu{
		Name: "zing", Age: 27,
	})

	log.Println(string(data))
	log.Println(strings.ReplaceAll(string(data), "{", ""))
}
