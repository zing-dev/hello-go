package main

import (
	"context"
	"fmt"
	"sync"
)

func HelloWorld() {
	fmt.Println("Hello, World!")
}

func HelloWorld2() {
	c := make(chan struct{})
	go func() {
		fmt.Println("Hello, World chan!")
		c <- struct{}{}
	}()
	<-c
}

func HelloWorld3() {
	var g sync.WaitGroup
	g.Add(1)
	go func() {
		fmt.Println("Hello, World WaitGroup!")
		g.Done()
	}()
	g.Wait()
}

func HelloWorld4() {
	cond := sync.NewCond(new(sync.Mutex))
	go func() {
		fmt.Println("Hello, World Cond!")
		cond.Signal()
	}()
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
}

func HelloWorld5() {
	ctx, cancel := context.WithCancel(context.Background())
	go cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Hello, World WithCancel!")
	}
}

func main() {
	HelloWorld()
	HelloWorld2()
	HelloWorld3()
	HelloWorld4()
	HelloWorld5()
}
