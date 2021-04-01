package main

import (
	"net/http"
	"sync"
	"testing"
	"time"
)

func testClose() {
	i := 0
	for i < 100 {
		_, _ = http.Get("http://192.168.0.251:7777/close?id=1")
		time.Sleep(time.Second)
		_, _ = http.Get("http://192.168.0.251:7777/close?id=2")
		time.Sleep(time.Second)
		i++
	}
}

func testRun() {
	i := 0
	for i < 100 {
		_, _ = http.Get("http://192.168.0.251:7777/run?id=1")
		time.Sleep(time.Second)
		_, _ = http.Get("http://192.168.0.251:7777/run?id=2")
		time.Sleep(time.Second)
		i++
	}
}

func TestStart(t *testing.T) {
	g := sync.WaitGroup{}
	g.Add(2)
	go func() {
		testRun()
		g.Done()
	}()
	go func() {
		testClose()
		g.Done()
	}()
	g.Wait()
}
