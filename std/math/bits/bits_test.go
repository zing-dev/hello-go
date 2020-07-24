package bits

import (
	"fmt"
	"math"
	"math/bits"
	"testing"
	"unsafe"
)

func TestUintSize(t *testing.T) {
	var x uint
	if want := unsafe.Sizeof(x) * 8; bits.UintSize != want {
		t.Fatalf("UintSize = %d; want %d", bits.UintSize, want)
	} else {
		t.Logf("UintSize = %d; want %d", bits.UintSize, want)
	}
}

func TestLeadingZeros(t *testing.T) {
	t.Log(bits.LeadingZeros(0))
	t.Log(bits.LeadingZeros(1))
	t.Log(bits.LeadingZeros(1 << 62))
	t.Log(bits.LeadingZeros(1<<63 - 1))
	t.Log(bits.LeadingZeros(1 << 63))
}

func TestLeadingZeros8(t *testing.T) {
	t.Log(math.MaxUint8) //256
	t.Log(bits.LeadingZeros8(1))
	t.Log(bits.LeadingZeros8(1 << 2))
	t.Log(bits.LeadingZeros8(1 << 4))
	t.Log(bits.LeadingZeros8(math.MaxUint8))
}

func TestLeadingZeros16(t *testing.T) {
	t.Log(math.MaxInt16) //32767
	t.Log(bits.LeadingZeros16(1))
	t.Log(bits.LeadingZeros16(1 << 2))
	t.Log(bits.LeadingZeros16(1 << 4))
	t.Log(bits.LeadingZeros16(math.MaxInt16))
	t.Log(bits.LeadingZeros16(math.MaxUint16))
}

func TestLeadingZeros64(t *testing.T) {
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1, bits.LeadingZeros64(1))
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1<<2, bits.LeadingZeros64(1<<2))
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1<<4, bits.LeadingZeros64(1<<4))
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1<<8, bits.LeadingZeros64(1<<8))
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1<<32, bits.LeadingZeros64(1<<32))
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1<<62, bits.LeadingZeros64(1<<62))
}

func TestTrailingZeros(t *testing.T) {
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 14, bits.TrailingZeros8(14))
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 1<<2, bits.TrailingZeros8(1<<2))
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 1<<3, bits.TrailingZeros8(1<<3))
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 1<<7, bits.TrailingZeros8(1<<7))
}

func TestOnesCount(t *testing.T) {
	fmt.Printf("OnesCount(%08b) = %d\n", 1<<2, bits.OnesCount(1<<2))
	fmt.Printf("OnesCount(%08b) = %d\n", 1<<3, bits.OnesCount(1<<3))
	fmt.Printf("OnesCount(%08b) = %d\n", 1<<7, bits.OnesCount(1<<7))
	fmt.Printf("OnesCount(%08b) = %d\n", 155, bits.OnesCount(155))
}

func TestRotateLeft8(t *testing.T) {
	fmt.Printf("RotateLeft8(%08b) = (%08b),%d\n", 1<<1, bits.RotateLeft8(1<<1, 1), bits.RotateLeft8(1<<1, 1))
	fmt.Printf("RotateLeft8(%08b) = (%08b),%d\n", 1<<2, bits.RotateLeft8(1<<2, 1), bits.RotateLeft8(1<<2, 1))
	fmt.Printf("RotateLeft8(%08b) = (%08b),%d\n", 1<<2, bits.RotateLeft8(1<<2, 2), bits.RotateLeft8(1<<2, 2))
	fmt.Println(1<<3, bits.RotateLeft(1, 3))
	fmt.Println(1 >> 5)
}

func TestRotateLeft(t *testing.T) {
	fmt.Printf("RotateLeft(%08b) = %d\n", 1<<10, bits.RotateLeft(1<<10, 1))
}

func TestReverse8(t *testing.T) {
	fmt.Printf("%d => Reverse8(%08b) = (%08b) <= %d\n", 1<<1, 1<<1, bits.Reverse8(1<<1), bits.Reverse8(1<<1))
	fmt.Printf("%d => Reverse8(%08b) = (%08b) <= %d\n", 1<<2, 1<<2, bits.Reverse8(1<<2), bits.Reverse8(1<<2))

}
