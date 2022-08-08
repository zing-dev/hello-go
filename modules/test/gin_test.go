package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//
//names[first]=thinkerou&names[second]=tianou

//ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
func TestMapAsQueryStringOrPostForm(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8181")
}
