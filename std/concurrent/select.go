package concurrent

import (
	"log"
	"time"
)

func select1() {
	//fatal error: all goroutines are asleep - deadlock!
	select {}
}

func select2() {
	select {
	default:
		log.Println("default")
	}
}

func select3() {
	var A = make(chan int)
	go func() {
		A <- 1
	}()

	select {
	case a := <-A:
		log.Println(a)
		//default:
		//	log.Println("default")
	}
}

func select4() {
	var A = make(chan string)
	go func() {
		for {
			time.Sleep(time.Second)
			A <- time.Now().Format(time.Kitchen)
		}
	}()
	for {
		select {
		case a := <-A:
			log.Println(a)
		default:
			log.Println("default")
			time.Sleep(time.Second)
		}
	}
}

func select5() {
	var A = make(chan string)
	var B = make(chan string)

	go func() {
		time.Sleep(time.Second)
		A <- time.Now().Format(time.Stamp)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		B <- time.Now().Format(time.Stamp)
	}()

	for {
		select {
		case a := <-A:
			log.Println(a)
		case b := <-B:
			log.Println(b)
		default:
			log.Println("default")
			time.Sleep(time.Second)
		}
	}
}

func select6() {
	ch := make(chan int)
	select {
	case <-ch:
	case <-time.After(time.Second * 1): // 利用time来实现，After代表多少时间后执行输出东西
		log.Println("超时啦!")
	}
}

func select7() {
	ch := make(chan int, 1) // 注意这里给的容量是1
	select {
	case ch <- 2:
		log.Println("22222")
	case <-ch:
		log.Println("<-")
	default:
		log.Println("通道channel已经满啦，塞不下东西了!")
	}
}
