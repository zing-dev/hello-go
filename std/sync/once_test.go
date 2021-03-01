package sync_test

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Load() {
	if arr == nil {
		arr = map[int]int{}
	}
}

func NewOnce() map[int]int {
	once.Do(Load)
	return arr
}

var (
	arr  map[int]int
	once sync.Once
)

func BenchmarkErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func(i int) {
			Load()
		}(i)
	}
}

func BenchmarkOnce(b *testing.B) {
	var g sync.WaitGroup
	for i := 0; i < b.N; i++ {
		g.Add(1)
		go func(i int) {
			NewOnce()
			g.Done()
		}(i)
	}
	g.Wait()
}

func TestName(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					t.Log("done")
					return // return结束该goroutine，防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		t.Log(n)
		if n == 5 {
			break
		}
	}
	cancel()
	time.Sleep(time.Millisecond)
}
