package binary

import (
	"encoding/binary"
	"log"
	"testing"
)

func TestInt16ToBytes(t *testing.T) {
	log.Println(Int16ToBytes(-1))
	log.Println(Int16ToBytes(1))
	log.Println(Int16ToBytes(1 << 2))
	log.Println(Int16ToBytes(1 << 4))
	log.Println(Int16ToBytes(1<<8 - 1))
	log.Println(Int16ToBytes(1 << 8))
	log.Println(Int16ToBytes(1 << 16))
}

func TestIntToBytes(t *testing.T) {
	log.Println(IntToBytes(1))
	log.Println(IntToBytes(1 << 2))
	log.Println(IntToBytes(1 << 4))
	log.Println(IntToBytes(1<<8 - 1))
	log.Println(IntToBytes(1 << 8))
	log.Println(IntToBytes(1 << 16))
}

func TestBytesToInt(t *testing.T) {
	log.Println(BytesToInt(IntToBytes(1 << 8)))
	log.Println(-1 << 7)
	log.Println(BytesToInt(IntToBytes(-1 << 7)))
	log.Println(BytesToInt(IntToBytes(-1 << 8)))
}

func TestIntegersToBytes(t *testing.T) {
	log.Println(IntegersToBytes([]int32{1, 2}))
	log.Println(BytesToIntegers(IntegersToBytes([]int32{1, 2})))
}

func TestSize(t *testing.T) {
	t.Log(binary.Size(int32(1)))
	t.Log(binary.Size(byte(1)))
	t.Log(binary.Size([10]byte{}))
	t.Log(binary.Size(struct{}{}))      //0
	t.Log(binary.Size(map[byte]byte{})) //-1
}
