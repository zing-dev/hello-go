package basic

import (
	"log"
	"testing"
)

func TestError(t *testing.T) {
	if e1, err := E1(true); err != nil {
		log.Println(err)
	} else if e2, err := E2(true); err != nil {
		log.Println(err)
	} else if e3, err := E3(true); err != nil {
		log.Println(err)
	} else {
		log.Println(e1, e2, e3)
	}
}
