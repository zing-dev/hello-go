//go:generate fileb0x b0x.yml
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	// your embedded files import here ...
	"fileb0x-demo/resources"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	e := echo.New()
	e.Debug = true

	// enable any filename to be loaded from in-memory file system
	e.GET("/*", echo.WrapHandler(resources.Handler))

	// read ufo.html from in-memory file system
	h, err := resources.ReadFile("world.html")
	if err != nil {
		log.Fatal(err)
	}

	// convert to string
	html := string(h)

	// serve ufo.html through "/"
	e.GET("/", func(c echo.Context) error {
		// serve it
		return c.HTML(http.StatusOK, html)
	})

	// try it -> http://localhost:1337/
	// http://localhost:1337/ufo.html
	// http://localhost:1337/public/README.md
	_ = open.Run("http://localhost:1337/")
	_ = e.Start(":1337")
}
