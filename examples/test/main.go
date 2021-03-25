package main

import "test/race"

type T interface {
	Run()
}

var App T = new(race.E3)

func main() {
	App.Run()
}
