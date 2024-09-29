package basic

import (
	"log"
	"testing"
	"time"
)

// 执行不到
func TestSelect(t *testing.T) {
	for {
		select {
		case <-time.After(time.Second * 9):
			t.Log("10秒")
			return
		case <-time.After(time.Second * 2):
			t.Log("2秒")
		case <-time.After(time.Second):
			t.Log("1秒")
		case <-time.After(time.Second / 2):
			t.Log("1/2秒")
		case <-time.After(time.Second / 5):
			t.Log("1/5秒")
		default:
			t.Log("default")
		case <-time.After(time.Second * 3):
			log.Println("sleep 3")

		case <-time.After(time.Second * 5):
			log.Println("sleep 5")
		}
	}
}

func TestSelectSleep(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-time.After(time.Second):
			log.Println("After")
		case <-ticker.C:
			log.Println("ticker")
		default:
			time.Sleep(time.Second * 2)
			log.Println("default")
		}
	}
}

// 可以执行
func TestSelectTicker(t *testing.T) {
	var (
		ticker001 = time.NewTicker(time.Second / 100)
		ticker01  = time.NewTicker(time.Second / 10)
		ticker1   = time.NewTicker(time.Second)
		ticker2   = time.NewTicker(time.Second * 2)
		ticker5   = time.NewTicker(time.Second * 5)
		ticker10  = time.NewTicker(time.Second * 10)
	)
	for {
		select {
		case <-ticker001.C:
			t.Log("1/100秒")
		case <-ticker01.C:
			t.Log("1/10秒")
		case <-ticker1.C:
			t.Log("1秒")
		case <-ticker2.C:
			t.Log("2秒")
		case <-ticker5.C:
			t.Log("5秒")
		case <-ticker10.C:
			t.Log("10秒")
			return
		}
	}
}
