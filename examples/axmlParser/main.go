package main

import (
	"github.com/lunny/axmlParser"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("run app.apk")
	}
	listener := new(axmlParser.AppNameListener)
	_, err := axmlParser.ParseApk(os.Args[1], listener)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listener.VersionName", listener.VersionName)
	log.Println("listener.VersionCode", listener.VersionCode)
	log.Println("listener.ActivityName", listener.ActivityName)
	log.Println("listener.PackageName", listener.PackageName)
}
