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
