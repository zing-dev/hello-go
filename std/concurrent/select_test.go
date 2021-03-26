package concurrent

import (
	"golang.org/x/tour/tree"
	"log"
	"testing"
	"time"
)

func TestSelect1(t *testing.T) {
	select1()
}

func TestSelect2(t *testing.T) {
	select2()
}

func TestSelect3(t *testing.T) {
	select3()
}

func TestSelect4(t *testing.T) {
	select4()
}

func TestSelect5(t *testing.T) {
	select5()
}

func TestSelect6(t *testing.T) {
	select6()
}

func TestSelect7(t *testing.T) {
	select7()
}

func TestSelect8(t *testing.T) {
	select8()
}

func TestSelect9(t *testing.T) {
	select9()
}

func TestSelect10(t *testing.T) {
	select10()
}

func TestSelect11(t *testing.T) {
	select11()
}

func TestSelect20(t *testing.T) {
	select20()
}

func TestSelect001(t *testing.T) {
	a := make(chan bool)
	time.AfterFunc(time.Second, func() {
		log.Println("start")
		select {
		case a <- true:
			log.Println("aaa")
		case _, ok := <-a:
			if !ok {
				a <- true
			}
		default:
			log.Println("default")
		}
	})
	go func() {
		for {
			log.Println("go start")
			select {
			case a <- true:
				log.Println("go aaa")
			case _, ok := <-a:
				if !ok {
					a <- true
				}
			default:
				log.Println("go default")
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case a := <-a:
			log.Println("over", a)
			return
		default:
			log.Println("sleep")
			time.Sleep(time.Second * 5)
		}
	}
}

func TestSelectTour4(t *testing.T) {
	c := make(chan int, 20)
	go func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}(cap(c), c)
	for i := range c {
		log.Println(i)
	}
}

func TestSelectTour5(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			log.Printf("%2d%3d", i+1, <-c)
		}
		quit <- 0
	}()
	func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				log.Println("quit")
				return
			}
		}
	}(c, quit)
}

func TestSelectTour6(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			log.Println("tick.")
		case <-boom:
			log.Println("BOOM!")
			return
		default:
			log.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		log.Println(x, y)
		if x != y {
			return false
		}
	}
	return true
}

func TestSelectTour7(t *testing.T) {
	log.Println("测试 Walk 函数。")
	ch := make(chan int)
	tt := tree.New(1)
	go Walk(tt, ch)
	for i := 0; i < 10; i++ {
		a := <-ch
		log.Println(a)
	}
	log.Println("测试 Same 函数。")
	log.Println(Same(tree.New(1), tree.New(1)))
	log.Println(Same(tree.New(1), tree.New(5)))
}
