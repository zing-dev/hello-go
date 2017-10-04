package main

import (
	"log"
	"os"
)


func main() {
	var (
		newFile *os.File
		err     error
	)

	newFile, err = os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
