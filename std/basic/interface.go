package basic

import (
	"encoding/json"
	"fmt"
	"log"
)

type Coder struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Languages []string `json:"languages"`
}

var (
	coder = Coder{
		Name:      "zing",
		Age:       26,
		Languages: []string{"Golang", "Java", "C/C++", "PHP"},
	}
)

func I1() {
	data, err := json.Marshal(coder)
	if err != nil {
		log.Fatal(err)
	}
	//{"name":"zing","age":26,"languages":["Golang","Java","C/C++","PHP"]}
	fmt.Println(string(data))
}

func I2() {
	var coder2 interface{}
	//coder2 = coder
	coder2 = &coder
	data, err := json.Marshal(coder2)
	if err != nil {
		log.Fatal(err)
	}
	//{"name":"zing","age":26,"languages":["Golang","Java","C/C++","PHP"]}
	fmt.Println(string(data))
}

func I3() {
	data, err := json.Marshal([]interface{}{
		coder, coder,
	})
	if err != nil {
		log.Fatal(err)
	}
	//[{"name":"zing","age":26,"languages":["Golang","Java","C/C++","PHP"]},{"name":"zing","age":26,"languages":["Golang","Java","C/C++","PHP"]}]
	fmt.Println(string(data))
}
