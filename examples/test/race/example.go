package race

import (
	"log"
	"sync"
	"time"
)

type E1 struct {
	v int
}

func (e E1) Run() {
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				e.v++
			}
		}()
	}
	time.Sleep(time.Second)
}

type E2 struct {
	v int
}

func (e E2) Run() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				e.v++
			}
			c <- e.v
			if i == 9 {
				close(c)
			}
		}(i)
	}
	for v := range c {
		log.Println(v)
	}
}

type E3 struct {
	v int
}

func (e E3) Run() {
	g := sync.WaitGroup{}
	c := make(chan int, 1)
	for i := 0; i < 10; i++ {
		g.Add(1)
		go func(i int) {
			defer g.Done()
			c <- 0
			for j := 0; j < 10; j++ {
				e.v++
			}
			<-c
		}(i)
	}
	g.Wait()
	log.Println(e.v)

}
