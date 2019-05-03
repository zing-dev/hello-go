package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	_ = c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	engine := gin.Default()
	engine.GET("/", formHandler)
	//engine.POST("/", formHandler)
	time.AfterFunc(time.Millisecond, func() {
		const host = " http://localhost:8080/ "
		const header = ` -H "Content-Type:application/json" `
		const post = ` -X POST `
		const data = ` --data '{"colors": ["red","yellow"]}' `
		fmt.Println(header + data + host)
		//command := exec.Command("curl", header,post,data,host)
		command := exec.Command("curl", host+"?colors[]=red")
		bytes, err := command.Output()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bytes))

	})
	_ = engine.Run()
}
