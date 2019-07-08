package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	token = ""
	host  = "http://192.168.10.10:8000"
	route = map[string]string{
		"login":  "/auth/login?username=%s&password=%s",
		"detail": "/auth/detail?token=%s",
	}
	client = http.Client{}
)

func init() {

}

func login() string {

	resp, err := client.Get(fmt.Sprintf("%s%s", host, fmt.Sprintf(route["login"], "zing", "zing")))

	defer func() {
		e := resp.Body.Close()
		if e != nil {
			log.Fatalln("close err")
		}
	}()

	if err != nil {
		log.Fatalln("get err")
	}

	if resp.StatusCode != 200 {
		log.Fatalln("StatusCode != 200")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("read err")
	}

	b := map[string]interface{}{}
	err = json.Unmarshal(body, &b)

	if err != nil {
		log.Fatal("unmarshal err")
	}

	data := b["data"].(map[string]interface{})
	token = data["token"].(string)
	return token
}

func detail() {

	resp, err := client.Get(fmt.Sprintf("%s%s", host, fmt.Sprintf(route["detail"], token)))

	if err != nil {
		log.Fatalln("get err")
	}

	if resp.StatusCode != 200 {
		log.Fatalln("StatusCode != 200")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("detail ReadAll err")
	}
	body := map[string]interface{}{}

	err = json.Unmarshal(bytes, &body)

	if err != nil {
		log.Fatalln("detail Unmarshal err")

	}
	data := body["data"].(map[string]interface{})
	fmt.Println(data["id"])
	fmt.Println(data["username"])
	fmt.Println(data["password"])

}

func main() {

	token = login()
	detail()

}
