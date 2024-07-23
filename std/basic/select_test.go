package basic

import (
	"log"
	"testing"
	"time"
)

// 执行不到
func TestSelect(t *testing.T) {
func TestSelectAfter(t *testing.T) {
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
func TestSelect(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	ticker3 := time.NewTicker(time.Second * 3)
	ticker5 := time.NewTicker(time.Second * 5)
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
		case <-ticker.C:
			log.Println("1s")
			time.Sleep(time.Second * 3)
		case <-ticker3.C:
			log.Println("3s")
		case <-ticker5.C:
			t.Log("5秒")
		case <-ticker10.C:
			t.Log("10秒")
			return
		}
	}
			log.Println("5s")
		}

	}
}
