package time

import (
	"fmt"
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

func TestNOW(t *testing.T) {
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
	time.Sleep(time.Second)
	log.Println(NOW())
}

func TestTime2(t *testing.T) {
	time2()
}

func TestTime3(t *testing.T) {
	time3()
}

func TestTime4(t *testing.T) {
	time4()
}

func TestTime5(t *testing.T) {
	time5()
}

func TestTime6(t *testing.T) {
	time6()
}

func TestTime7(t *testing.T) {
	time7()
}

func TestRound(t *testing.T) {
	round()
}

func TestRound2(t *testing.T) {
	round2()
}

func interval(interval time.Duration) {
	if interval != time.Minute && interval != time.Minute*5 && interval != time.Minute*15 {
		log.Println("err")
	} else {
		log.Println("ok")
	}
}

func TestT1(t *testing.T) {
	var i = 10
	log.Println(time.Minute == time.Duration(i)*time.Second)
	log.Println(time.Minute == time.Duration(i)*time.Second*6)
	log.Println(time.Minute != time.Duration(i)*time.Second*6)
	log.Println(time.Minute != time.Duration(i)*time.Second*6 || time.Second == time.Minute/time.Duration(i*6))
	interval(time.Minute)
	interval(time.Minute + 1)
	log.Println(time.Now().Truncate(time.Minute * 4))
}

func TestTime(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Second)
	log.Println(time.Now().Sub(start))
}

func TestTime8(t *testing.T) {
	i := 10
	i2 := time.Second * 10
	log.Println(i, i2)
	log.Println(i, i2.Seconds())
	log.Println(i, i2.Minutes())
	log.Println(i, i2.Milliseconds())
	log.Println(time.Duration(i), i2.Milliseconds())
	log.Println(time.Duration(i).Nanoseconds(), i2.Nanoseconds())
	log.Println(time.Duration(i).Seconds(), i2)
	log.Println(time.Duration(i) * time.Second)
}

func TestTime9(t *testing.T) {
	log.Println(time.Now().Unix())
	log.Println(time.Now().UnixNano())
	log.Println(time.Now().UnixNano() / 1000000)
	log.Println(fmt.Sprintf("%v", time.Now().UnixNano()/1000000))
	log.Println(fmt.Sprintf("\\/Date(%13d+0800)\\/", time.Now().UnixNano()/1000000))
}
