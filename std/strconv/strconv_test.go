package strconv

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"
)

func TestFormatFloat(t *testing.T) {
	log.Println(strconv.FormatFloat(1.1, 'g', 10, 32))
	log.Println(strconv.FormatFloat(1.11111, 'g', 10, 32))
	log.Println(strconv.FormatFloat(116.397128, 'g', 10, 32))
	log.Println(strconv.FormatFloat(116.397128, 'f', 10, 32))
	log.Println(fmt.Sprintf("%f", 116.397128))
	log.Println(fmt.Sprintf("%f", 116.397128))
	log.Println(float32(math.Ceil(1.001002)))
	log.Println(float32(math.Round(1.001002)))
	log.Println(float32(math.Floor(1.001002)))
	log.Println(strconv.ParseFloat("10.0010101", 32))
	log.Println(strconv.ParseFloat(fmt.Sprintf("%.2f", 10.1010101), 32))
}
