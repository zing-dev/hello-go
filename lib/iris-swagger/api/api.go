package api

import (
	"fmt"
	"iris-swagger/web"
	"net/http"

	"github.com/kataras/iris/v12"
)

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.Message "We need ID!!"
// @Failure 404 {object} web.Message "Can not find ID"
// @Router /test/{id} [get]
func GetStringByInt(ctx iris.Context) {
	err := web.Message{}
	fmt.Println(err)
	ctx.StatusCode(200)
	_, _ = ctx.WriteString("hello")
}

//
// @Summary 快速上手接口测试
// @Description 快速上手接口测试,这是我的描述
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "id"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "id错误"
// @Failure 404 {object} web.Message "id没找到"
// @Router /quick/{id} [get]
func TestQuick(ctx iris.Context) {
	i, err := ctx.Params().GetInt("id")
	fmt.Println(i)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	if i < 0 {
		ctx.StatusCode(http.StatusBadRequest)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusBadRequest,
			Message: "id错误",
		})
		return
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: "成功",
	})
}

type ParamSearch struct {
	Id   int    `json:"id,omitempty" format:"int"`
	Name string `json:"name,omitempty" format:"string"`
}

// @Summary 快速上手接口 POST 测试
// @Description 快速上手接口测试,这是我的描述
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "id"
// @Param   id     body    ParamSearch     true        "id"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "id错误"
// @Failure 404 {object} web.Message "id没找到"
// @Router /quick-post/{id} [post]
func TestQuickPost(ctx iris.Context) {
	var p = &ParamSearch{}
	err := ctx.ReadJSON(p)
	fmt.Println(p)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "param没找到",
		})
		return
	}
	i, err := ctx.Params().GetInt("id")
	fmt.Println(ctx.PostValue("id"))
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	if i < 0 {
		ctx.StatusCode(http.StatusBadRequest)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusBadRequest,
			Message: "id错误",
		})
		return
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: "成功",
	})
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.Message "We need ID!!"
// @Failure 404 {object} web.Message "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(ctx iris.Context) {}

type Pet3 struct {
	ID int `json:"id"`
}
