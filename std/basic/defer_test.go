package basic

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	t.Log("start....")
	n := rand.Intn(10)
	if n > 1 {
		t.Log("n > 1")
		return
	}

	defer func() {
		t.Log("defer end...")
	}()
	t.Log(" end...")
}

func BenchmarkDefer(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		TestDefer2(&testing.T{})
	}
	b.StopTimer()
}

func TestDefer2(t *testing.T) {
	var (
		flagCh      = make(chan int)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*9)
		start       = time.Now()
	)
	defer func() {
		cancel()
		close(flagCh)
		t.Log("defer end...")
	}()
	go func() {
		for {
			n := rand.Intn(100)
			if n > 80 {
				flagCh <- n
				return
			}
			time.Sleep(time.Second)
		}
	}()

	select {
	case flag := <-flagCh:
		t.Logf("flag: %v", flag)
	case <-ctx.Done():
		t.Logf("ctx: %v", ctx.Err())
	}

	t.Logf("==> %s", time.Since(start))
}
