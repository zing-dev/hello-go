package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 {
		dir += "/" + os.Args[1]
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.FileServer(http.Dir(dir)).ServeHTTP(writer, request)
	})
	fmt.Println("=============start===========")
	_ = http.ListenAndServe("localhost:8080", nil)
}
