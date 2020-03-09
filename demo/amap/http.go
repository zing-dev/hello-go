package amap

import (
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")

	})
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "09_08_10_08__09_08_32_40.json")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
