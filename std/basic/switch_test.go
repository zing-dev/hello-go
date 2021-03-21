package basic

import (
	"log"
	"testing"
)

func TestSwitch(t *testing.T) {
	a := 1
	switch a {
	case 1, 2:
		log.Println("ok")
	case 3:
		log.Println("3")
	default:
		log.Println("default")
	}

}
