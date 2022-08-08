package main

import (
	"github.com/lunny/axmlParser"
	"log"
	"os"
	"testing"
)

func TestAXmlParserQuick(t *testing.T) {
	dir, err := os.Getwd()
	log.Println(dir, err)
}

func TestAXmlParser(t *testing.T) {
	path := "../../../../flutter/otdr_online/build/app/outputs/flutter-apk/app-release.apk"
	listener := new(axmlParser.AppNameListener)
	apk, err := axmlParser.ParseApk(path, listener)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("apk.Namespaces: ", apk.Namespaces)
	log.Println("apk.ResCount: ", apk.ResCount)
	log.Println("apk.ResourcesIds: ", apk.ResourcesIds)
	log.Println("listener.VersionName", listener.VersionName)
	log.Println("listener.VersionCode", listener.VersionCode)
	log.Println("listener.ActivityName", listener.ActivityName)
	log.Println("listener.PackageName", listener.PackageName)
}
