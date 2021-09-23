package api

import (
	"fmt"
	"iris-swagger/web"
	"net/http"

	"github.com/kataras/iris/v12"
)

// GetStringByInt
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

// TestQuick
// @Summary 快速上手接口测试
// @Description 快速上手接口测试,这是我的描述
// @Tags 接口
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

// TestQuickPost
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

// TestQuickPostForm
// @Summary 快速上手接口 POST Form测试
// @Description 快速上手接口测试,这是我的描述
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id     path    int     true        "id"
// @Param   id     formData    int     true        "id"
// @Param   name   formData    string     true     "名字"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "id错误"
// @Failure 404 {object} web.Message "id没找到"
// @Router /quick-post-form/{id} [post]
func TestQuickPostForm(ctx iris.Context) {
	id, err := ctx.PostValueInt("id")
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	name := ctx.PostValue("name")
	id2, err := ctx.Params().GetInt("id")
	fmt.Println(ctx.PostValue("id"))
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	if id2 < 0 {
		ctx.StatusCode(http.StatusBadRequest)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusBadRequest,
			Message: "id错误",
		})
		return
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: fmt.Sprintf("成功: path: id: %d; form: id:%d, name: %s", id2, id, name),
	})
}

// TestQuickPostQuery
//@Summary 快速上手接口 POST Query测试
// @Description 快速上手接口测试,这是我的描述
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id     path    int     true        "id"
// @Param   id     query    int     true        "id"
// @Param   name   query    string     true     "名字"
// @Param   id2     formData    int     true        "id2"
// @Param   name2   formData    string     true     "名字2"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "id错误"
// @Failure 404 {object} web.Message "id没找到"
// @Router /quick-post-query/{id} [post]
func TestQuickPostQuery(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	name := ctx.URLParam("name")
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	id2, err := ctx.PostValueInt("id2")
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "id没找到",
		})
		return
	}
	name2 := ctx.PostValue("name2")
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: fmt.Sprintf("成功: query: id: %d, name: %s;form: id: %d, name: %s", id, name, id2, name2),
	})
}

// TestQuickFile
// @Summary 快速上手接口 POST File测试
// @Description 快速上手接口测试,这是我的描述
// @Accept  mpfd
// @Produce  json
// @Param   file   formData    file     true     "上传文件"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "file没找到"
// @Router /quick-file [post]
func TestQuickFile(ctx iris.Context) {
	_, header, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.JSON(web.Message{
			Code:    http.StatusNotFound,
			Message: "file没找到",
		})
		return
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: fmt.Sprintf("成功:form: name: %s,size: %d", header.Filename, header.Size),
	})
}

// GetStructArrayByString
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
