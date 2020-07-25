package math

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"testing"
)

func TestDecimal(t *testing.T) {
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
