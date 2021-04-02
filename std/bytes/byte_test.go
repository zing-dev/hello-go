package bytes

import (
	"bytes"
	"compress/bzip2"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
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
	log.Println(string(rune(67)))
	log.Println(string(rune(65)))
	log.Println(string(rune(65)) + string(rune(65)))
	log.Println(string(rune(65 + 65)))
	log.Println('A')
	log.Println('Z')
	log.Println(91 - 65)
	log.Println(string(rune(91)))
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
	log.Println(string(rune(65)))      //A  65  + 0
	log.Println(string(rune(66)))      //B  65  + 1
	log.Println(string(rune(90)))      //Z  65  + 25
	log.Println(string(rune(91)))      //AA 65  + 26 + 0 //
	log.Println(string(rune(92)))      //AB 65  + 26 + 1
	log.Println(string(rune(92 + 26))) //AB 65  + 26 + 1
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
	str := []string{"hello", "world"}
	arr, _ := json.Marshal(str)
	log.Println(string(arr))

	log.Println(strings.Join(str, ""))
}

func TestBz2(t *testing.T) {
	data, err := ioutil.ReadFile("C:\\Go\\src\\compress\\bzip2\\testdata\\Isaac.Newton-Opticks.txt.bz2")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(data))
	reader := bzip2.NewReader(bytes.NewReader(data))
	if data, err := ioutil.ReadAll(reader); err != nil {
		log.Fatal(err)
	} else {
		log.Println(len(data))
	}
}
