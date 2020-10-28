package concurrent

import (
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
