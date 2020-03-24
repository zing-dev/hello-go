package math

import "math"

func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) / 1e2
}
