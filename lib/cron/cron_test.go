package cron

import (
	"github.com/robfig/cron/v3"
	"log"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	id1, err := c.AddFunc("*/10 * * * * *", func() {
		log.Println("hello world")
	})
	log.Println(id1)
	if err != nil {
		log.Fatal(err)
	}
	id2, err := c.AddFunc("*/20 * * * * *", func() {
		log.Println("hello cron...")
	})
	log.Println(id2)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()

	id3, err := c.AddFunc("*/5 * * * * *", func() {
		log.Println("hello cron /5...")
	})
	log.Println(id3)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case <-time.After(time.Minute / 2):
		log.Println("-------")
		c.Remove(id1)
	}
}
