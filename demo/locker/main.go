package main

import (
	"fmt"
	"time"
)

type Data struct {
	Date time.Time
	Ok   *bool
	Msg  string
}

var data = new(Data)

func write(ok bool, msg string) {
	data.Date = time.Now()
	data.Ok = &ok
	data.Msg = msg
}

func read() {
	fmt.Printf("%s %v %s\n", data.Date.Format("15:04:05"), *data.Ok, data.Msg)
}

func main() {
	write(true, "init")
	go func() {
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				go write(true, "true")
			} else {
				go write(false, "false")
			}
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			go read()
		}
	}()
	time.Sleep(time.Second * 2)
}
