package context

import (
	"context"
	"fmt"
	"log"
	"sync"
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
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	go func() {
		time.Sleep(time.Second * 30)
		cancel()
		log.Println("parent cancel")
	}()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	go func() {
		time.Sleep(time.Second * 20)
		defer cancel()
		log.Println("child cancel")
	}()
	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("overslept")
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				log.Println("parent already cancel")
			}
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("child already timeout")
			}
			time.Sleep(time.Second / 10)
			return
		}
	}
}

func TestTimeout(t *testing.T) {
	type App struct {
		ctx    context.Context
		cancel context.CancelFunc
	}

	type Child struct {
		ctx    context.Context
		cancel context.CancelFunc
	}
	ctx, cancel := context.WithCancel(context.Background())
	app := App{ctx: ctx, cancel: cancel}

	ctx, cancel = context.WithTimeout(app.ctx, time.Second*30)
	child := Child{ctx: ctx, cancel: cancel}
	go func() {
		time.Sleep(time.Second * 30)
		child.cancel()
	}()
	ticket := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticket.C:
			log.Println("timeout 5 second")
		case <-time.After(time.Second * 3):
			log.Println("go...")
			ticket.Reset(time.Second * 5)
			time.Sleep(time.Second * 6)
		case <-app.ctx.Done():
			log.Println("app done")
		case <-child.ctx.Done():
			if child.ctx.Err() == context.Canceled {
				log.Println("cancel..")
			}
			if child.ctx.Err() == context.DeadlineExceeded {
				log.Println("deadline..")
			}
			return
		}
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

func TestName111(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("do something...")
		case <-ctx.Done():
			fmt.Println("main", ctx.Err())
			return
		}
	}
}

func product(channel chan<- int, i int) {
	for {
		fmt.Println("product", i)
		channel <- i
		time.Sleep(time.Second * time.Duration(1))
		i--
		if i == 0 {
			close(channel)
			fmt.Println("product over")
			return
		}
	}
}

func customer(channel <-chan int, quick chan struct{}) {
	for {
		message, ok := <-channel
		if ok {
			fmt.Println("customer", message)
		} else {
			fmt.Println("customer over")
			quick <- struct{}{}
			return
		}
	}
}

func TestName112(t *testing.T) {
	channel := make(chan int, 5)
	quick := make(chan struct{})
	go product(channel, 10)
	go customer(channel, quick)
	<-quick
}

var once sync.Once

type manager struct{ name string }

var single *manager

func Singleton() *manager {
	once.Do(func() {
		single = &manager{"a"}
	})
	return single
}
