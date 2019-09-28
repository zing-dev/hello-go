package binary_test

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

func TestLittleEndian(t *testing.T) {

	var data [4]byte
	binary.LittleEndian.PutUint16(data[0:], math.MaxInt16)
	binary.LittleEndian.PutUint16(data[2:], math.MaxInt16-math.MaxInt8)
	fmt.Printf("% x\n", data)

	u := binary.LittleEndian.Uint16(data[0:])
	fmt.Println(u)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt16 - math.MaxInt8)

	fmt.Println(binary.LittleEndian.Uint16(data[2:]))
	fmt.Printf("% x", binary.LittleEndian.Uint16(data[2:]))
}
