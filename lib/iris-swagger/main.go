package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"iris-swagger/api"
	_ "iris-swagger/docs"
)

// @title swagger 接口开发测试
// @version v0.0.1
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @tag.name 接口
// @tag.description 这是接口文档

// @tag.name 接口2
// @tag.description 这是接口文档2

// @tag.name CMS
// @tag.description CMS接口文档
//go:generate swag init
//
// @host localhost:8080
// @BasePath /api/v1
func main() {
	app := iris.New()
	app.Get("/api/v1/quick/{id}", api.TestQuick)
	app.Post("/api/v1/quick-post/{id}", api.TestQuickPost)
	app.Post("/api/v1/quick-post-form/{id}", api.TestQuickPostForm)
	app.Post("/api/v1/quick-post-query/{id}", api.TestQuickPostQuery)
	app.Post("/api/v1/quick-file", api.TestQuickFile)

	user := new(api.User)
	{
		app.Post("/api/v1/user/login", user.Login)
		app.Get("/api/v1/user/list", user.List)
	}

	app.Get("/api/v1/test/{id}", api.GetStringByInt)

	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler))
	app.Run(iris.Addr(":8080"))
}
