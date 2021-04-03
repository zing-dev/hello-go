package main

import (
	"test/net"
)

type T interface {
	Run()
}

var App T = new(net.App)

func main() {
	App.Run()
}
