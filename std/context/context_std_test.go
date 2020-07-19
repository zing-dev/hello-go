package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithCancel(t *testing.T) {
	c1, cancel := context.WithCancel(context.Background())

	if got, want := fmt.Sprint(c1), "context.Background.WithCancel"; got != want {
		t.Errorf("c1.String() = %q want %q", got, want)
	}

	o := otherContext{c1}
	c2, _ := context.WithCancel(o)

	c4, cancel4 := context.WithCancel(context.Background())

	contexts := []context.Context{c1, o, c2, c4}

	for i, c := range contexts {
		if d := c.Done(); d == nil {
			t.Errorf("c[%d].Done() == %v want non-nil", i, d)
		}
		if e := c.Err(); e != nil {
			t.Errorf("c[%d].Err() == %v want nil", i, e)
		}

		select {
		case x := <-c.Done():
			t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
		default:
		}
	}

	cancel()
	time.Sleep(100 * time.Millisecond) // let cancellation propagate

	for i, c := range contexts {
		select {
		case <-c.Done():
			t.Logf("<-c[%d].Done() %s", i, c.Err())
		default:
			t.Errorf("<-c[%d].Done() blocked, but shouldn't have", i)
		}
		if e := c.Err(); e != context.Canceled {
			t.Errorf("c[%d].Err() == %v want %v", i, e, context.Canceled)
		}
	}
	cancel4()
	select {
	case <-c4.Done():
		t.Logf("<-c[4].Done() %s", c4.Err())
	default:
		t.Errorf("<-c[%d].Done() blocked, but shouldn't have", 4)
	}
}
