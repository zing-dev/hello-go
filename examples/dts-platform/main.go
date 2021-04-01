package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	app := App{}
	go app.Run()
	s := <-c
	// Block until a signal is received.
	app.Close()
	fmt.Println("Got signal:", s)
}
