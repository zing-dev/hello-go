package basic

import (
	"errors"
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

func TestErrorCase(t *testing.T) {
	var (
		e1 = errors.New("error1")
		e2 = errors.New("error2")
		e3 = errors.New("error3")
	)

	switch errors.New("error1") {
	case e1:
		log.Println(e1)
	case e2:
		log.Println(e2)
	case e3:
		log.Println(e3)
	default:
		log.Println("fuck")
	}
	switch errors.Unwrap(errors.New("error1")) {
	case e1:
		log.Println(e1)
	case e2:
		log.Println(e2)
	case e3:
		log.Println(e3)
	default:
		log.Println("fuck")
	}
	log.Println(errors.Is(errors.New("error1"), e1))
	log.Println(errors.Is(e1, e1))
	log.Println(errors.Is(e1, errors.New("error1")))
}
