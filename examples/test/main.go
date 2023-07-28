package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"test/dts-client"
)

type T interface {
	Run()
}

var App T = new(dts.Client)

func main() {
	//App.Run()

	dir, err := os.Getwd()
	log.Println(dir, err)
	log.Println(execPath())
}

func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(file), err
}
