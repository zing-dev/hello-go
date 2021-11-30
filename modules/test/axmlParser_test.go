package test

import (
	"github.com/lunny/axmlParser"
	"log"
	"testing"
)

func TestAXmlParser(t *testing.T) {
	listener := new(axmlParser.AppNameListener)
	apk, err := axmlParser.ParseApk("C:\\Users\\admin\\OneDrive\\workspace\\flutter\\otdr_online\\build\\app\\outputs\\apk\\release\\app-release.apk", listener)
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
