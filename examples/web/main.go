package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	port := 8080
	if len(os.Args) == 3 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal("port error")
		} else {
			port = p
		}
		dir += "/" + os.Args[2]
	} else {
		log.Fatal("error: web port dir need!")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.FileServer(http.Dir(dir)).ServeHTTP(writer, request)
	})
	fmt.Println(fmt.Sprintf("=============start at :%d===========", port))
	err = http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
