package binary

import (
	"bytes"
	"encoding/binary"
	"log"
)

func Int16ToBytes(i int) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

func IntToBytes(i int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func IntegersToBytes(i []int32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, i)
	if err != nil {
		log.Println(err)
		return nil
	}
	return buf.Bytes()
}

func BytesToIntegers(b []byte) []int32 {
	i := make([]int32, len(b)/4)
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, i)
	if err != nil {
		log.Println(err)
		return nil
	}
	return i
}

func BytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}

func BytesToInt16(b []byte) int {
	return int(binary.BigEndian.Uint16(b))
}
