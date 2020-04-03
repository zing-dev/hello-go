package math

import (
	"log"
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
}
