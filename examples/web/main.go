package main

import (
	"fmt"
	"github.com/judwhite/go-svc"
	"log"
	"net/http"
	"os"
	"strconv"
)

// implements svc.Service
type program struct{}

func main() {
	prg := program{}

	// call svc.Run to start your program/service
	// svc.Run will call Init, Start, and Stop
	if err := svc.Run(&prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())

	// write to "example.log" when running as a Windows Service
	if env.IsWindowsService() {
	}

	return nil
}

func (p *program) Start() error {
	log.Printf("Starting...\n")
	go run()
	return nil
}

func (p *program) Stop() error {
	log.Printf("Stopping...\n")
	log.Printf("Stopped.\n")
	return nil
}

func run() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	port := 9999
	if len(os.Args) == 3 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal("port error")
		} else {
			port = p
		}
		dir += "/" + os.Args[2]
	} else if len(os.Args) == 1 {
		dir += "/."
	} else {
		log.Fatal("error: web port dir need!")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("from ", request.RemoteAddr)
		http.FileServer(http.Dir(dir)).ServeHTTP(writer, request)
	})
	fmt.Println(fmt.Sprintf("=============start at :%d===========", port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
