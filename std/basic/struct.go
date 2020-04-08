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
