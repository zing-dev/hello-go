package basic

import (
	"log"
	"testing"
	"time"
)

func TestSelectAfter(t *testing.T) {
	for {
		select {
		case <-time.After(time.Second * 3):
			log.Println("sleep 3")

		case <-time.After(time.Second * 5):
			log.Println("sleep 5")
		}
	}
}

func TestSelect(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	ticker3 := time.NewTicker(time.Second * 3)
	ticker5 := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Println("1s")
			time.Sleep(time.Second * 3)
		case <-ticker3.C:
			log.Println("3s")
		case <-ticker5.C:
			log.Println("5s")
		}

	}
}
