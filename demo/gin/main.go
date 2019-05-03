package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Group("/json/", func(context *gin.Context) {
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		log.Fatal("启动gin服务失败")
	}
}
