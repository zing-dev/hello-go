package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"
)

func TestWriteFloat64(t *testing.T) {
	var pi = 3.141592653589793
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%x\n", buf.Bytes())
}

func TestWriteStruct(t *testing.T) {
	zing := People{
		Id:   1,
		Name: "zing",
		Age:  25,
	}
	data, _ := json.Marshal(zing)
	pack := Package{
		Id:     12,
		Length: len(data),
		Data:   data,
	}
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.LittleEndian, pack.Id)
	err = binary.Write(&buf, binary.LittleEndian, pack.Length)
	err = binary.Write(&buf, binary.LittleEndian, pack.Data)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	b := buf.Bytes()
	fmt.Printf("%d\n", b[0])
	data = make([]byte, 10)
	buf2 := bytes.NewReader(buf.Bytes())
	err = binary.Read(buf2, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(len(data))
	fmt.Println(data[0])
	fmt.Println(data[1])
}
