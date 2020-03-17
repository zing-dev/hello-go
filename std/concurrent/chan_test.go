package concurrent

import (
	"log"
	"testing"
)

func TestChannel(t *testing.T) {
	notify := make(chan int)
	log.Println(notify)
	log.Println(notify == nil)
	close(notify)
	log.Println(notify == nil)
	notify = make(chan int)
	log.Println(notify)
	close(notify)
	log.Println(notify)
	log.Println(len(notify))
}

func TestChannel1(t *testing.T) {
	channel1()
}

func TestChannel2(t *testing.T) {
	channel2()
}

func TestChannel3(t *testing.T) {
	channel3()
}

func TestChannel4(t *testing.T) {
	channel3()
}
