package basic

import (
	"encoding/json"
	"fmt"
	"log"
)

func S1() {
	data, err := json.Marshal(coder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	log.Println(string(data))

	data, err = json.MarshalIndent(coder, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	log.Println(string(data))

	/*var coder2 *Coder
	err = json.Unmarshal(data, coder2)
	if err != nil {
		//Unmarshal:json: Unmarshal(non-pointer basic.Coder)
		log.Fatal("Unmarshal:",err)
	}
	fmt.Println(coder2)*/

	/*var coder3 Coder
	err = json.Unmarshal(data, coder3)
	if err != nil {
		//Unmarshal:json: Unmarshal(non-pointer basic.Coder)
		log.Fatal("Unmarshal:",err)
	}
	fmt.Println(coder3)*/

	var coder4 Coder
	err = json.Unmarshal(data, &coder4)
	if err != nil {
		log.Fatal("Unmarshal:", err)
	}
	fmt.Println(coder4)
	fmt.Println(&coder4)
}

func S2() {
	type C struct {
		C string `json:"c"`
	}
	type B struct {
		B1 string  `json:"b1"`
		B2 int     `json:"b2"`
		B3 float64 `json:"b3"`
		C  C       `json:"c"`
	}
	type A struct {
		A string `json:"a"`
		B B      `json:"b"`
	}

	a := A{
		A: "A",
		B: B{
			B1: "B1",
			B2: 1,
			B3: 1.2,
			C: C{
				C: "C",
			},
		},
	}
	data, _ := json.MarshalIndent(a, "", "    ")
	fmt.Println(string(data))

	var face interface{}
	var face2 map[string]interface{}
	_ = json.Unmarshal(data, &face)
	_ = json.Unmarshal(data, &face2)
	fmt.Println(face)
	fmt.Println(face2)
	fmt.Println(face2["b"].(map[string]interface{})["c"])

}
