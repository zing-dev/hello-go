package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

// wonderful 1 ==> go:embed $HOME/path/file
// wonderful 2 ==> go:embed $(os.Getwd())/path/file

//go:embed main.go
var str string

func stringTest() {
	log.Println("================")
	fmt.Println(str)
	log.Println("================")
}

//go:embed main.go
var b []byte

func byteTest() {
	log.Println("================")
	fmt.Println(b)
	log.Println("================")
}

//go:embed main.go
var f embed.FS

func fileTest() {
	content, err := f.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("================")
	fmt.Println(string(content))
	log.Println("================")
}

//go:embed *.go *.mod
var dir embed.FS

func matchTest() {
	file, err := dir.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(info.Name())

	file, err = dir.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	info, err = file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(info.Name())

	strings, err := fs.Glob(dir, "*.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(strings)
}
