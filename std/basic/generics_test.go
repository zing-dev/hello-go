package basic

import (
	"fmt"
	"testing"
	"time"
)

import "github.com/samber/lo"

func TestFind(t *testing.T) {
	people, ok := lo.Find[People]([]People{{Id: 1, Name: "zing"}, {Id: 2}}, func(p People) bool {
		return p.Id == 1
	})
	fmt.Println(people, ok)
}

func TestChan(t *testing.T) {
	var (
		send = make(chan int, 10)
	)
	go func() {
		i := 0
		for {
			send <- i
			i++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			s := <-send
			fmt.Println("1-> ", s)
		}
	}()
	for {
		s := <-send
		fmt.Println("2-> ", s)
	}
}

func TestName31(t *testing.T) {
	printFooOrBar(foo{"foo"})
	printFooOrBar(bar{"bar"})

	d1 := []foo{{"1"}, {"22"}, {"333"}, {"4444"}, {"55555"}}
	d2 := []any{foo{"1"}, foo{"22"}, bar{"333"}, bar{"4444"}, bar{"55555"}, bar{"666666"}, foo{"7777777"}}
	fmt.Println(find(d1))
	fmt.Println(find(d2))
	fmt.Println(filter(d1))
	fmt.Println(filter(d2))
}
