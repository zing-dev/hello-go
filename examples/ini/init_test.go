package ini

import (
	"gopkg.in/ini.v1"
	"log"
	"testing"
)

type IP struct {
	Value []string `ini:"value,omitempty,allowshadow"`
}

func TestIni(t *testing.T) {
	file, err := ini.ShadowLoad("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	ip := &IP{}
	err = file.MapTo(&ip)
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Section("Ip").ReflectFrom(&ip); err != nil {
		log.Fatal(err)
	}
	if err := file.SaveTo("config.ini"); err != nil {
		log.Fatal(err)
	}
	log.Println(ip)
}
