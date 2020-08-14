package rand

import (
	"math/rand"
	"testing"
)

func TestNewSource(t *testing.T) {
	source := rand.NewSource(1)
	t.Log(source.Int63())

	source = rand.NewSource(1)
	t.Log(source.Int63())

	source = rand.NewSource(2)
	t.Log(source.Int63())
}

func TestRand(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	t.Log(r.Int63())
	t.Log(r.Int())
	t.Log(r.Int31())

	t.Log(r.Uint32())
	t.Log(r.Uint64())

	t.Log(r.Int31n(1))
	t.Log(r.Int31n(2))
	t.Log(r.Int31n(3))

	t.Log(r.Float32())
	t.Log(r.Float64())
}
