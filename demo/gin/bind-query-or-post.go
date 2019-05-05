package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Name     string    `form:"name" json:"name"`
	Address  string    `form:"address" json:"address"`
	Birthday time.Time `form:"birthday"json:"birthday" time_format:"2006-01-02 15:04:05" time_utc:"1"`
}

func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		user := &User{}
		err := context.ShouldBind(user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
				"status":  false,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message": "success",
				"status":  true,
				"data":    user,
			})
		}
	})
	_ = engine.Run()
}

/**
data: {name: "zing", address: "wuxi", birthday: "1994-09-18T15:04:05Z"}
message: "success"
status: true
*/

//$ curl -X GET "localhost:8080?name=zing&address=wuxi&birthday=1994-09-18"
