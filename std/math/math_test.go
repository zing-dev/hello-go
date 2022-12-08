package math

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"testing"
)

func TestDecimal(t *testing.T) {
	log.Println(Decimal32(1.001, 1))
	log.Println(Decimal32(1.51, 1))
	log.Println(Decimal32(1.91001, 1))
	log.Println(Decimal32(0.019, 1))
	log.Println(Decimal32(0.001, 1))
	log.Println(Decimal32(0.005, 1))
	log.Println(Decimal32(0.055, 1))
	log.Println(Decimal32(0.555, 1))
	log.Println(Decimal32(100.005, 1))
	log.Println(Decimal32(202.60000610351562, 1))
	log.Println(Decimal32(202.699999, 1))
	log.Println(Decimal32(202.699999, 1) * 2.001)
	v := Decimal32(202.123546, 1)
	log.Println(fmt.Sprintf("%v %f %f", v, v, 1.1))

	log.Println(Decimal(1.001))
	log.Println(Decimal(1.51))
	log.Println(Decimal(1.91001))
	log.Println(Decimal(0.019))
	log.Println(Decimal(0.001))
	log.Println(Decimal(0.005))
	log.Println(Decimal(0.055))
	log.Println(Decimal(0.555))
	log.Println(Decimal(100.005))
}

func TestMath(t *testing.T) {
	Math()
	max := math.Max(1, 1)
	log.Println(max)
}

func TestAbs(t *testing.T) {
	fmt.Println(math.Abs(-1.1))
	fmt.Println(math.Abs(1))
	fmt.Println(math.MaxFloat64 > math.MaxUint64)
	fmt.Println(bits.Len64(math.Float64bits(math.MaxFloat64)))

}

func TestUint8(t *testing.T) {
	t.Log(math.MaxUint8)
	t.Log(math.MaxUint8 / 60)
	t.Log(math.MaxUint8 / 60 / 60)
}

func TestUint16(t *testing.T) {
	t.Log(math.MaxUint16)
	t.Log(math.MaxUint16 / 60)
	t.Log(math.MaxUint16 / 60 / 60)
}

func TestDecimal2(t *testing.T) {
	log.Println(Decimal2(1.11112))
	log.Println(Decimal2(1.22223333))
	log.Println(Decimal3(1.22223333, 1))
	log.Println(Decimal3(1.22223333, 2))
	log.Println(Decimal3(1.22223333, 3))
	log.Println(Decimal3(1.22223333, 4))
	v := Decimal3(157.39999389648438, 1)
	log.Println(v)
	log.Println(v * 2.00)
	log.Println(Decimal32(157.39999389648438, 1))
}
