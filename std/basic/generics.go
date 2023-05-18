package basic

import (
	"fmt"
	"github.com/samber/lo"
)

type foo struct {
	foo string
}

type bar struct {
	bar string
}

func printFooOrBar[T foo | bar](v T) {
	switch v := (any)(v).(type) {
	case foo:
		fmt.Println(v.foo)
	case bar:
		fmt.Println(v.bar)
	}
}

func find[T foo | bar | any](v []T) (T, bool) {
	return lo.Find[T](v, func(v T) bool {
		if f, ok := (any)(v).(foo); ok {
			return len(f.foo) > 30
		}
		if f, ok := (any)(v).(bar); ok {
			return len(f.bar) > 5
		}
		return false
	})
}

func filter[T foo | bar | any](v []T) []T {
	return lo.Filter[T](v, func(v T, _ int) bool {
		if f, ok := (any)(v).(foo); ok {
			return len(f.foo) > 3
		}
		if f, ok := (any)(v).(bar); ok {
			return len(f.bar) > 5
		}
		return false
	})
}

type I[T any] interface {
	query() I[T]
}

type IS struct {
	test string
}

type IS2 struct {
	test string
}

func (i IS) query() I[IS] {
	return i
}

func (i IS2) query() I[IS] {
	return i
}

type C[T I[T]] struct {
	c T
}

func (c C[T]) check() I[T] {
	return c.c.query()
}

func (c C[T]) get() T {
	return any(c.c).(T)
}

func (c C[T]) getRef() *T {
	t2 := any(c.c).(T)
	return &t2
}
