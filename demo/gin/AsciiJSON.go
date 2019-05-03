package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		context.AsciiJSON(http.StatusOK, gin.H{
			"lang":     "GO语言",
			"examples": "AsciiJSON",
			"tag":      "<br>",
		})
	})
	_ = engine.Run()
}
