package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:generate fileb0x b0x.yml

func handle(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		FileHandle(writer, request)
	case http.MethodPost:
		FileUpload(writer, request)
	default:
	}
}

func main() {
	flag.Parse()
	config = Config{
		Port: flag.Arg(0),
		Dir:  flag.Arg(1),
	}
	if config.Port == "" {
		config.Port = "8081"
	}
	if config.Dir == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			config.Dir = "./"
		} else {
			config.Dir = dir + "/wf-uploads"
		}
	}

	_, err := os.Stat(config.Dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(config.Dir, 0x777)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println(fmt.Sprintf("uploads dir %s and port %s", config.Dir, config.Port))
	http.HandleFunc("/", handle)
	log.Println(fmt.Sprintf("server start from :%s...", config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)
	log.Println("server end: ", err)
}
