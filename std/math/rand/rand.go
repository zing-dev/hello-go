package rand

import (
	"math"
	"math/rand"
	"time"
)

func rand1(f float32) float32 {
	arr := []float32{-1, 1}
	rand.Seed(time.Now().Unix())
	if f < 1 {
		f += arr[rand.Int()%2] * rand.Float32() / 5
	} else if f < 50 {
		f += arr[(rand.Int()+1)%2] * rand.Float32() * 3
	} else {
		f += arr[rand.Int()%2] * rand.Float32() * 5
	}
	return float32(math.Trunc(float64(f*1e2+0.5)) / 1e2)
}
