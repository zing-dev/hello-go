package main

import (
	"fmt"
	"github.com/rjeczalik/notify"
	"log"
	"os"
	"path/filepath"
	"time"
)

var WaitFor = func(path string, e notify.Event, timeout time.Duration) bool {
	dir, file := filepath.Split(path)
	c := make(chan notify.EventInfo, 1)

	if err := notify.Watch(dir, c, e); err != nil {
		log.Fatal(err)
	}
	// Clean up watchpoint associated with c. If Stop was not called upon
	// return the channel would be leaked as notify holds the only reference
	// to it and does not release it on its own.
	defer notify.Stop(c)

	t := time.After(timeout)

	for {
		select {
		case ei := <-c:
			fmt.Println(filepath.Base(ei.Path()), file, ei.Event())
			if filepath.Base(ei.Path()) == file {
				return true
			}
		case <-t:
			return false
		}
	}
}

var filename = "test.txt"

func main() {
	go func() {
		os.Remove(filename)
		time.Sleep(time.Second * 3)
		os.Create(filename)
		time.Sleep(time.Second * 3)
		if err := os.Remove(filename); err != nil {
			fmt.Printf("remove: %s", err)
		}
	}()

	go func() {
		if WaitFor(filename, notify.Create, time.Second*10) {
			log.Println(filename, "locked")
		}
	}()

	go func() {
		if WaitFor(filename, notify.All, time.Second*10) {
			log.Println(filename, "locked2")
		}
	}()

	if WaitFor(filename, notify.Remove, time.Second*15) {
		log.Println(filename, "unlocked")
	}

	fmt.Println("over")
}
