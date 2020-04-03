package math

import (
	"log"
	"math"
)

func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) / 1e2
}

func Math() {
	log.Println("Abs", math.Abs(-1))
	log.Println("Ceil", math.Ceil(11.11))
	log.Println("Floor", math.Floor(11.61))
	log.Println("Cbrt", math.Cbrt(11.61))
	log.Println("Dim", math.Dim(6.1, 12.2))
	log.Println("NaN", math.NaN())
	log.Println("Trunc", math.Trunc(12.1))
	log.Println("Sqrt", math.Sqrt(16))
	log.Println("Round", math.Round(16.1))
	log.Println("Round", math.Round(16.5))

	log.Println("RoundToEven", math.RoundToEven(16.1))
	log.Println("RoundToEven", math.RoundToEven(16.5))
	log.Println("RoundToEven", math.RoundToEven(16.6))
	log.Println("RoundToEven", math.RoundToEven(16.7))
	log.Println("RoundToEven", math.RoundToEven(16.8))

	log.Println("Mod", math.Mod(11, 12))
}
