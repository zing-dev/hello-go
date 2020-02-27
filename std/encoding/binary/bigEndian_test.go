package binary_test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBigEndian(t *testing.T) {

	fmt.Println(binary.BigEndian.String())
	fmt.Println(binary.BigEndian.GoString())

	b := make([]byte, 4)
	u := binary.BigEndian.Uint16([]byte{1, 2})
	fmt.Println(u) //258
	binary.BigEndian.PutUint16(b[0:], u)
	fmt.Printf("% x\n", b) //

	binary.BigEndian.PutUint16(b[0:], 0x03e8)
	binary.BigEndian.PutUint16(b[2:], 0x07d0)
	fmt.Printf("% x\n", b)
	fmt.Printf("%x\n", b)
	fmt.Printf("%b\n", b)
	fmt.Printf("%d\n", 0x03e8)                  //1000
	fmt.Printf("%d\n", 0x07d0)                  //2000
	fmt.Println(binary.BigEndian.Uint16(b[0:])) //1000
	fmt.Println(binary.BigEndian.Uint16(b[2:])) //2000

	b = make([]byte, 8)
	binary.BigEndian.PutUint32(b[0:], 0x03e803e8)
	binary.BigEndian.PutUint32(b[4:], 0x07d007d0)
	fmt.Printf("% x\n", b) //03 e8 03 e8 07 d0 07 d0

}
