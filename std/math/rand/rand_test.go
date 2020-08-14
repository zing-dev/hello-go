package rand

import (
	"log"
	"testing"
)

func TestRand1(t *testing.T) {
	log.Println(rand1(0.4))
	log.Println(rand1(25))
	log.Println(rand1(88))
}
