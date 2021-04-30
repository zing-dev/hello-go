package main

import (
	"test/dts-client"
)

type T interface {
	Run()
}

var App T = new(dts.Client)

func main() {
	App.Run()
}
