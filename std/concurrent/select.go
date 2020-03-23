package concurrent

import (
	"fmt"
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

func select8() {
	i := 0
	defer fmt.Println("for循环外")

	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				return
			}
		}
		fmt.Println("for循环内 i=", i)
	}
}

func select9() {
	i := 0
Loop:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				break Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}

	fmt.Println("for循环外")
}

func select10() {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}
Loop:
	fmt.Println("for循环外")
}

func select11() {
	t1 := time.NewTicker(time.Second * 5)
	t2 := time.NewTicker(time.Second * 5)
	defer func() {
		t1.Stop()
		t2.Stop()
	}()
	for {
		select {
		case <-t1.C:
			log.Println("========t1==start============")
			T1()
			log.Println("========t1==end============")
		case <-t2.C:
			log.Println("========t2===start===========")
			T2()
			log.Println("========t2===end===========")
		}
	}
}

func T1() {
	i := 0
	for i < 1000 {
		if i%50 == 0 {
			log.Println(i)
		}
		i++
	}
}
func T2() {
	i := 0
	for i < 2000 {
		if i%80 == 0 {
			log.Println(i)
		}
		i++
	}
}
