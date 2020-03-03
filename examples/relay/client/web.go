package main

import (
	"log"
	"net/http"
	"time"
)

func NewWeb() {
	c := NewClient()
	go c.Heart()
	http.HandleFunc("/toggle", func(writer http.ResponseWriter, request *http.Request) {
		time.AfterFunc(time.Second, func() {
			_ = c.RelayOpenAll()
			time.Sleep(time.Second)
			_ = c.RelayCloseAll()
		})
		_, _ = writer.Write([]byte("toggle ok....."))
	})
	http.HandleFunc("/c", func(writer http.ResponseWriter, request *http.Request) {
		c.Close()
		_, _ = writer.Write([]byte("close ok....."))
	})
	http.HandleFunc("/open", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(c.RelayOpenAll().Error()))
	})
	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(c.RelayCloseAll().Error()))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
