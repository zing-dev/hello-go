package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	go time.AfterFunc(time.Nanosecond, func() {
		const host = "http://localhost:8080/"
		command := exec.Command("curl", host+"getb?field_a=hello&field_b=world")
		bytes, _ := command.Output()
		fmt.Println(string(bytes))

		command = exec.Command("curl", host+"getc?field_a=hello&field_c=world")
		bytes, _ = command.Output()
		fmt.Println(string(bytes))

		command = exec.Command("curl", host+"getd?field_x=hello&field_d=world")
		bytes, _ = command.Output()
		fmt.Println(string(bytes))

		command = exec.Command("curl", host+"getd?field_=hello&field_d=world")
		bytes, _ = command.Output()
		fmt.Println(string(bytes))
	})
	_ = r.Run()

}
