package byte

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	byte1()
}

func TestName2(t *testing.T) {
	type A struct{}
}

func TestName3(t *testing.T) {
	log.Println(76)
	log.Println(string(67))
	log.Println(string(65))
	log.Println(string(65) + string(65))
	log.Println(string(65 + 65))
	log.Println('A')
	log.Println('Z')
	log.Println(91 - 65)
	log.Println(string(91))
	log.Println((91 - 65) / 26)
	log.Println((91 - 65) % 26)
	log.Println((91 + 26 - 65) / 26)
	log.Println((91 + 26 - 65) % 26)
}

func TestName4(t *testing.T) {
	log.Println(row(65))
	log.Println(row(66))
	log.Println(row(90))
	log.Println(row(91))
	log.Println(row(92))
	log.Println(row(93))
	log.Println(row(93))
}

func TestName5(t *testing.T) {
	log.Println(string(65))      //A  65  + 0
	log.Println(string(66))      //B  65  + 1
	log.Println(string(90))      //Z  65  + 25
	log.Println(string(91))      //AA 65  + 26 + 0 //
	log.Println(string(92))      //AB 65  + 26 + 1
	log.Println(string(92 + 26)) //AB 65  + 26 + 1
}

func row(i int) string {
	str := ""
	i = i - 65
	for i != 0 {
		b := i % 26
		str = string(b+65-1) + str
		i = i / 26
	}
	return str
}
