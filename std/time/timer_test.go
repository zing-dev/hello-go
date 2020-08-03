package time

import (
	"log"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 3)
	log.Println("before 3 second")
	<-timer.C
	log.Println("after 3 second")
}

func TestStop(t *testing.T) {
	log.Println("-------")
	timer := time.NewTimer(time.Second * 3)
	log.Println("-------")
	timer.Stop()
	log.Println("-------")

}

func TestReset(t *testing.T) {
	timer := time.NewTimer(time.Second * 3)
	if !timer.Stop() {
		log.Println("run")
		<-timer.C
	}
	timer.Reset(time.Second * 5)
	log.Println("reset")
}

func TestAfter(t *testing.T) {
	log.Println("-------")
	<-time.After(time.Second * 3)
	log.Println("-------")
}

func TestAfterFunc(t *testing.T) {
	log.Println("-------------")
	time.AfterFunc(time.Second, func() {
		log.Println("-------------")
	})
	time.Sleep(time.Second * 3)
	log.Println("-------------")
}
