package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

//go:generate fileb0x b0x.yml

type config struct {
	port string
	dir  string
}

func handle(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		FileHandle(writer, request)
	case http.MethodPost:
		_, _ = writer.Write([]byte("post"))
	default:
	}
}

func main() {
	flag.Parse()
	c := config{
		port: flag.Arg(0),
		dir:  flag.Arg(1),
	}
	if c.port == "" {
		c.port = "8081"
	}
	if c.dir == "" {
		c.dir = "./"
	}
	log.Println(fmt.Sprintf("www dir %s, port %s", c.dir, c.port))
	http.HandleFunc("/", handle)
	fmt.Println(fmt.Sprintf("server start from :%s...", c.port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", c.port), nil)
	fmt.Println("server end: ", err)

}
