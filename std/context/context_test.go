package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBackground(t *testing.T) {
	c := context.Background()
	if c == nil {
		t.Fatalf("Background returned nil")
	}
	select {
	case x := <-c.Done():
		t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}
	if got, want := fmt.Sprint(c), "context.Background"; got != want {
		t.Errorf("Background().String() = %q want %q", got, want)
	} else {
		t.Log(got, want)
	}
}

func TestTODO(t *testing.T) {
	c := context.TODO()
	if c == nil {
		t.Fatalf("TODO returned nil")
	}
	select {
	case x := <-c.Done():
		t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}
	if got, want := fmt.Sprint(c), "context.TODO"; got != want {
		t.Errorf("TODO().String() = %q want %q", got, want)
	}
}

func TestWithCancel2(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					t.Log(n)
					n++
				default:
					t.Log("=> ", n)
					time.Sleep(time.Second / 2)
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
	time.Sleep(time.Second * 5)
}

func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-time.After(51 * time.Millisecond):
		fmt.Println("hello")
	case <-ctx.Done():
		fmt.Println("=>", ctx.Err())
	}
}

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func TestWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, "color")
}

func TestName(t *testing.T) {
	messages := make(chan int, 10)
	done := make(chan bool)
	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
