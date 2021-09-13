package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"iris-swagger/web"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age"`
	Address  string `json:"address,omitempty"`
}

type Users []*User

var (
	users = Users{
		{Id: 1, Name: "zing", Age: 25, Password: "zing", Address: "WuXi"},
		{Id: 2, Name: "trump", Age: 75, Password: "maga", Address: "USA"},
		{Id: 3, Name: "特朗普2", Age: 76, Address: "USA"},
		{Id: 4, Name: "特朗普3", Age: 77, Address: "USA"},
		{Id: 5, Name: "特朗普4", Age: 78, Address: "USA"},
	}
)

// @Summary 登录
// @Description 登录
// @Tags CMS
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name   formData    string     true     "用户名"
// @Param   password   formData    string     true     "密码"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "登录失败"
// @Router /user/login [post]
func (u *User) Login(ctx iris.Context) {
	name := ctx.PostValue("name")
	password := ctx.PostValue("password")
	for _, v := range users {
		if name == v.Name || password == v.Password {
			_, _ = ctx.JSON(web.Message{
				Code:    0,
				Message: fmt.Sprintf("成功:form: name: %s,password: %s", name, password),
			})
			return
		}
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: "登录失败",
	})
}

// @Summary 用户列表
// @Description 用户列表
// @Tags CMS
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   page   query    int     false     "第几页"
// @Param   size   query    int     false     "每页数据"
// @Success 200 {string} string	"成功"
// @Failure 400 {object} web.Message "登录失败"
// @Security ApiKeyAuth
// @Security BasicAuth
// @Router /user/list [get]
func (u *User) List(ctx iris.Context) {
	p := &Page{Total: len(users)}
	p.Check(ctx)
	p.Data = users[p.Start:p.End]
	if p.Error != nil {
		return
	}
	_, _ = ctx.JSON(web.Message{
		Code:    0,
		Message: "success",
		Data:    p,
	})
}
