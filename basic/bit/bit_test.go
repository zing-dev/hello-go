package bit

import (
	"encoding/binary"
	"log"
	"math"
	"testing"
)

func C(a interface{}) (data []byte) {
	switch a.(type) {
	case byte:
		b := a.(byte)
		data = make([]byte, 8)
		for k := range data {
			data[k] = (b & (1 << k)) >> k
		}
	case int16:
		i := a.(int16)
		d := make([]byte, 2)
		binary.LittleEndian.PutUint16(d, uint16(i))
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
	case uint16:
		i := a.(uint16)
		d := make([]byte, 2)
		binary.LittleEndian.PutUint16(d, i)
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
	case int32:
		i := a.(int32)
		d := make([]byte, 4)
		binary.LittleEndian.PutUint32(d, uint32(i))
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
		data = append(data, C(d[2])...)
		data = append(data, C(d[3])...)
	case uint32:
		i := a.(uint32)
		d := make([]byte, 4)
		binary.LittleEndian.PutUint32(d, i)
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
		data = append(data, C(d[2])...)
		data = append(data, C(d[3])...)
	case int64:
		i := a.(int64)
		d := make([]byte, 8)
		binary.LittleEndian.PutUint64(d, uint64(i))
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
		data = append(data, C(d[2])...)
		data = append(data, C(d[3])...)
		data = append(data, C(d[4])...)
		data = append(data, C(d[5])...)
		data = append(data, C(d[6])...)
		data = append(data, C(d[7])...)
	case uint64:
		i := a.(uint64)
		d := make([]byte, 8)
		binary.LittleEndian.PutUint64(d, i)
		data = append(data, C(d[0])...)
		data = append(data, C(d[1])...)
		data = append(data, C(d[2])...)
		data = append(data, C(d[3])...)
		data = append(data, C(d[4])...)
		data = append(data, C(d[5])...)
		data = append(data, C(d[6])...)
		data = append(data, C(d[7])...)
	}
	return data
}

func TestName(t *testing.T) {
	log.Println(C(byte(1))) //0 0 0 0 0 0 0 1
	log.Println(C(byte(2)))
	log.Println(C(byte(3)))
	log.Println(C(byte(255)))

	log.Println(C(int16(10)))
	log.Println(C(int16(16)))
	log.Println(C(int16(math.MaxInt16)))
	log.Println(C(uint16(math.MaxUint16)))
	log.Println(C(int32(math.MaxInt32)))
	log.Println(C(int64(math.MaxInt64)))
	log.Println(C(uint64(math.MaxUint64)))
}
