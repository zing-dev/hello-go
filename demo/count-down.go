package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr: "192.168.1.114:8888",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RemoteAddr)
		writer.Header().Add("Refresh", "3")

		in := time.Now()
		out := time.Date(in.Year(), in.Month(), in.Day(), 17, 0, 0, 0, time.Local)
		diff := out.Sub(in)
		_, _ = writer.Write([]byte("<h1>" + diff.String() + "</h1>"))
	})
	_ = server.ListenAndServe()

}
