package main

import (
	"log"
	"net/http"
)

func NewWeb() {
	c := NewClient()
	go c.Heart()
	http.HandleFunc("/open", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(c.Open().Error()))
	})
	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(c.Close().Error()))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
