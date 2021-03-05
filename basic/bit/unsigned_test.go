package bit

import (
	"log"
	"math"
	"testing"
)

//负数 -> 反码 -> 补码
func byte2Int8(b byte) int8 {
	return int8(b)
}

func int82Byte(i int8) byte {
	return byte(i)
}

func Test_int82Byte(t *testing.T) {
	b1 := byte(1)
	log.Printf("%4d %4x %9.8b \n", math.MinInt8, math.MinInt8, math.MinInt8)
	log.Printf("%4d %4x %9.8b \n", int82Byte(1), int82Byte(1), int82Byte(1))
	log.Printf("%4d %4x %9.8b \n", ^int82Byte(1), ^int82Byte(1), ^int82Byte(1))
	log.Printf("%4d %4x %9.8b \n", ^int82Byte(1)+1, ^int82Byte(1)+1, ^int82Byte(1)+1)
	log.Printf("%4d %4x %9.8b \n", int82Byte(-1), int82Byte(-1), int82Byte(-1))
	log.Printf("%4d %4x %9.8b 1\n", b1, b1, b1)
	log.Printf("%4d %4x %9.8b 1反\n", ^b1, ^b1, ^b1)
	log.Printf("%4d %4x %9.8b 负1\n", ^b1+b1, ^b1+b1, ^b1+b1)
	log.Printf("%4d %4x %9.8b 负1\n", int8(^b1+b1), int8(^b1+b1), int8(^b1+b1))
	log.Printf("%4d %4x %9.8b 负1\n", byte2Int8(uint8(math.MaxUint8)), byte2Int8(uint8(math.MaxUint8)), byte2Int8(uint8(math.MaxUint8)))
	log.Printf("%4d %4x %9.8b -1\n", -b1, -b1, -b1)

	log.Printf("%4d %4x %9.8b \n", int82Byte(math.MinInt8), int82Byte(math.MinInt8), int82Byte(math.MinInt8))
	log.Printf("%4d %4x %9.8b \n", int82Byte(math.MaxInt8), int82Byte(math.MaxInt8), int82Byte(math.MaxInt8))
}

func Test_byte2Int8(t *testing.T) {
	log.Printf("MinInt8  %d %x %b \n", math.MinInt8, math.MinInt8, math.MinInt8)
	log.Printf("MaxInt8  %d %x %b \n", math.MaxInt8, math.MaxInt8, math.MaxInt8)
	log.Printf("MaxUint8 %d %x %b \n", math.MaxUint8, math.MaxUint8, math.MaxUint8)
	log.Printf("%d %x %b \n", byte2Int8(math.MaxInt8), byte2Int8(math.MaxInt8), byte2Int8(math.MaxInt8))
	log.Printf("%d %x %b \n", byte2Int8(math.MaxUint8), byte2Int8(math.MaxUint8), byte2Int8(math.MaxUint8))
	log.Printf("%d %x %b \n", -math.MaxInt8, -math.MaxInt8, -math.MaxInt8)
}
