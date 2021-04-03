package net

import "os"

type App struct{}

func (a App) Run() {
	if len(os.Args) == 2 {
		if os.Args[1] == "s" {
			new(Server).Run()
		} else {
			new(Client).Run()
		}
	}
}
