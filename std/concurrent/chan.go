package concurrent

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type T int

func channel1() {
	notify := make(chan int)
	data := make(chan int, 100)
	go func() {
		<-notify
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100; i++ {
			data <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(data)

	}()
	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	notify <- 1
	time.Sleep(1 * time.Second)
	fmt.Println("通知1秒后接收到数据通道数据 ")
	for {
		if i, ok := <-data; ok {
			fmt.Println(i)
		} else {
			fmt.Println("接收不到数据中止循环")
			break
		}
	}
}

//1个发送者 N个接收者
func channel2() {
	notify := make(chan int)
	data := make(chan int, 100)
	go func() {
		<-notify
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100000; i++ {
			data <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(data)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	notify <- 1
	time.Sleep(1 * time.Second)
	fmt.Println("3秒后接受到数据通道数据 此时datach 在接收端已经关闭")
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				if j, ok := <-data; ok {
					fmt.Println(i, j)
				} else {
					break
				}
			}
		}(i)
	}
	time.Sleep(5 * time.Second)
}

//N 个发送者 1个接收者
//添加一个 停止通知 接收端告诉发送端不要发送了
func channel3() {
	dataCh := make(chan T, 1)
	stopCh := make(chan T)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			defer log.Println("协程", i, "over")
			for {
				value := T(rand.Intn(10000))
				select {
				case <-stopCh:
					fmt.Println(i, "接收到停止发送的信号")
					return
				case dataCh <- value:
				}
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("1秒后开始接收数据")
	for {
		if d, ok := <-dataCh; ok {
			fmt.Println(d)
			if d == 9999 {
				fmt.Println("当在接收端接收到9999时告诉发送端不要发送了")
				close(stopCh)
				return
			}
		} else {
			break
		}
	}
}

//N个发送者 M个接收者
func channel4() {
	dataCh := make(chan T, 100)
	toStop := make(chan string)
	stopCh := make(chan T)
	//简约版调度器
	go func() {
		if t, ok := <-toStop; ok {
			log.Println(t)
			close(stopCh)
		}
	}()
	//生产者
	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				id := strconv.Itoa(i)
				value := T(rand.Intn(10000))
				if value == 9999 {
					select {
					case toStop <- "sender# id:" + id + "to close":
					default:
					}
				}
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(i)
	}
	//消费者
	for i := 0; i < 20; i++ {
		go func(i int) {
			id := strconv.Itoa(i)
			for {
				select {
				case <-stopCh:
					return
				default:
				}
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == 9998 {
						select {
						case toStop <- "receiver# id:" + id + "to close":
						default:
						}
					}
					log.Println("receiver value :", value)
				}
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
