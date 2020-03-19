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

func TestName5(t *testing.T) {
	log.Println(string(65))      //A  65  + 0
	log.Println(string(66))      //B  65  + 1
	log.Println(string(90))      //Z  65  + 25
	log.Println(string(91))      //AA 65  + 26 + 0 //
	log.Println(string(92))      //AB 65  + 26 + 1
	log.Println(string(92 + 26)) //AB 65  + 26 + 1
}

func TestName4(t *testing.T) {
	log.Println(row(65))
	log.Println(row(66))
	log.Println(row(67))
	log.Println(row(90))
	log.Println(row(91))
	log.Println(row(92))
	log.Println(row(93))
	log.Println(row(100))
	log.Println(row(65 + 26 + 26))
}

func row(column int) string {
	column -= 64
	str := ""
	for column > 26 {
		i := column % 26
		if i == 0 {
			str = string(64+26) + str
			column = (column - 26) / 26
		} else {
			str = string(64+i) + str
			column = (column - i) / 26
		}
	}
	return string(column+64) + str
}
