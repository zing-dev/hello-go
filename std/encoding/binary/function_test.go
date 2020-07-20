package binary

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

func TestEncoding(t *testing.T) {
	if res, err := decodeByte(bytes.NewBuffer([]byte{0x56})); res != 0x56 || err != nil {
		t.Errorf("decodeByte([0x56]) did not return (0x56, nil) but (0x%X, %v)", res, err)
	}
	if res, err := decodeUint16(bytes.NewBuffer([]byte{0x56, 0x78})); res != 22136 || err != nil {
		t.Errorf("decodeUint16([0x5678]) did not return (22136, nil) but (%d, %v)", res, err)
	}
	if res := encodeUint16(22136); !bytes.Equal(res, []byte{0x56, 0x78}) {
		t.Errorf("encodeUint16(22136) did not return [0x5678] but [0x%X]", res)
	}

	strings := map[string][]byte{
		"foo":         {0x00, 0x03, 'f', 'o', 'o'},
		"\U0000FEFF":  {0x00, 0x03, 0xEF, 0xBB, 0xBF},
		"A\U0002A6D4": {0x00, 0x05, 'A', 0xF0, 0xAA, 0x9B, 0x94},
	}
	for str, encoded := range strings {
		if res, err := decodeString(bytes.NewBuffer(encoded)); res != str || err != nil {
			t.Errorf("decodeString(%v) did not return (%q, nil), but (%q, %v)", encoded, str, res, err)
		}
		if res := encodeString(str); !bytes.Equal(res, encoded) {
			t.Errorf("encodeString(%q) did not return [0x%X], but [0x%X]", str, encoded, res)
		}
	}

	lengths := map[int][]byte{
		0:         {0x00},
		127:       {0x7F},
		128:       {0x80, 0x01},
		16383:     {0xFF, 0x7F},
		16384:     {0x80, 0x80, 0x01},
		2097151:   {0xFF, 0xFF, 0x7F},
		2097152:   {0x80, 0x80, 0x80, 0x01},
		268435455: {0xFF, 0xFF, 0xFF, 0x7F},
	}
	for length, encoded := range lengths {
		if res, err := decodeLength(bytes.NewBuffer(encoded)); res != length || err != nil {
			t.Errorf("decodeLength([0x%X]) did not return (%d, nil) but (%d, %v)", encoded, length, res, err)
		}
		if res := encodeLength(length); !bytes.Equal(res, encoded) {
			t.Errorf("encodeLength(%d) did not return [0x%X], but [0x%X]", length, encoded, res)
		}
	}
}

func TestEncodingPackage(t *testing.T) {
	data := []byte("hello world")
	p := Package{
		Id:     127,
		Length: len(data),
		Data:   data,
	}
	buf := bytes.Buffer{}
	buf.WriteByte(p.Id)
	buf.Write(encodeUint16(uint16(p.Length)))
	buf.Write(encodeBytes(p.Data))

	reader := bytes.NewReader(buf.Bytes())
	id, err := decodeByte(reader)
	log.Println("id:", id, err)
	length, err := decodeUint16(reader)
	log.Println("length:", length, err)
	data, err = decodeBytes(reader)
	log.Println("data:", string(data), err)

	data, _ = json.Marshal(People{
		Id:   123,
		Name: "zing",
		Age:  25,
	})
	p = Package{
		Id:     127,
		Length: len(data),
		Data:   data,
	}
	buf = bytes.Buffer{}
	buf.WriteByte(p.Id)
	buf.Write(encodeUint16(uint16(p.Length)))
	buf.Write(encodeBytes(p.Data))

	reader = bytes.NewReader(buf.Bytes())
	id, err = decodeByte(reader)
	log.Println("id:", id, err)
	length, err = decodeUint16(reader)
	log.Println("length:", length, err)
	data, err = decodeBytes(reader)
	log.Println("data:", string(data), err)
}
