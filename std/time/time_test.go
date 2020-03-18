package time

import (
	"log"
	"testing"
	"time"
)

const (
	LocalDateTimeFormat = "2006-01-02 15:04:05"
)

var NOW = func() string {
	return time.Now().Format(LocalDateTimeFormat)
}

func TestTime1(t *testing.T) {
	time1()
}

func TestName(t *testing.T) {
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
}
