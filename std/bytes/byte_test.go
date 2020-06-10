package byte

import (
	"encoding/binary"
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
	log.Println(row(67))
	log.Println(row(90))
	log.Println(row(91))
	log.Println(row(92))
	log.Println(row(93))
	log.Println(row(100))
	log.Println(row(65 + 26 + 26))
}

func TestName5(t *testing.T) {
	log.Println(string(65))      //A  65  + 0
	log.Println(string(66))      //B  65  + 1
	log.Println(string(90))      //Z  65  + 25
	log.Println(string(91))      //AA 65  + 26 + 0 //
	log.Println(string(92))      //AB 65  + 26 + 1
	log.Println(string(92 + 26)) //AB 65  + 26 + 1
}

func TestName6(t *testing.T) {
	str := "a15297c10-9f1f-0605-5c54-4f41150d6e2e"
	log.Println([]byte(str))
	log.Println([]byte(str[:4]))
	log.Println(string([]byte{49, 53, 50}))
	log.Println([]byte("1"))
	log.Println([]byte("哈"))
	log.Println(string([]byte{229, 147, 136}))
	log.Println(binary.BigEndian.Uint16([]byte(str[:2])))
	log.Println(binary.BigEndian.Uint16([]byte(str[2:4])))
	a := make([]byte, 2)
	binary.BigEndian.PutUint16(a, 24881)
	log.Println(a)
}
func TestName7(t *testing.T) {
	data := make([]byte, 10)
	binary.BigEndian.PutUint16(data, 10)
	log.Println(data)
	binary.BigEndian.PutUint16(data, 100)
	log.Println(data)
	binary.BigEndian.PutUint16(data[2:4], 100)
	log.Println(data)
	binary.BigEndian.PutUint16(data[4:8], 256)
	log.Println(data)
	log.Println(append(data, 1))
}

func TestName8(t *testing.T) {
	byte2()
	//byte3()
}
