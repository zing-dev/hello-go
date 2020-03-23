package concurrent

import (
	"log"
	"time"
)

func run(t *time.Ticker, over chan bool) {
	t = time.NewTicker(time.Second * 2)
	defer func() {
		t.Stop()
		log.Println("over")
	}()
	for {
		select {
		case <-t.C:
			log.Println("here")
		case <-over:
			return
		}
	}
}

func select20() {
	var t *time.Ticker
	over := make(chan bool)
	go run(t, over)
	time.Sleep(time.Second * 10)
	log.Println("即将停止.....")
	//t.Stop()
	over <- false
	log.Println("已经停止.....")
	time.Sleep(time.Second * 3)
	go run(t, over)
	time.Sleep(time.Second * 10)
}
