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

func TestGenerics1(t *testing.T) {
	printFooOrBar(foo{"foo"})
	printFooOrBar(bar{"bar"})

	d1 := []foo{{"1"}, {"22"}, {"333"}, {"4444"}, {"55555"}}
	d2 := []any{foo{"1"}, foo{"22"}, bar{"333"}, bar{"4444"}, bar{"55555"}, bar{"666666"}, foo{"7777777"}}
	fmt.Println(find(d1))
	fmt.Println(find(d2))
	fmt.Println(filter(d1))
	fmt.Println(filter(d2))
}

func TestGenerics2(t *testing.T) {
	is := IS{test: "hello"}
	c := C[IS]{c: is}
	fmt.Println(c.check().(IS).test)
	fmt.Println(c.get().test)
	fmt.Println(c.getRef().test)
	//fmt.Println(any(1).(string))
	//fmt.Println(any(byte(1)).(string))
	//fmt.Println(interface{}("wocao").([]byte))
	//fmt.Println(interface{}("wocao").([]byte))
	fmt.Println(([]byte)("wocao"))
}
