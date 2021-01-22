package api

import (
	"github.com/kataras/iris/v12"
)

type Data struct {
	Pages int         `json:"pages"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type Page struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Pages int         `json:"pages"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Error error       `json:"-"`
	Start int         `json:"-"`
	End   int         `json:"-"`
}

func (p *Page) Check(ctx iris.Context) *Page {
	if ctx.URLParamExists("page") {
		p.Page, p.Error = ctx.URLParamInt("page")
		if p.Error != nil {
			return p
		}
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if ctx.URLParamExists("size") {
		p.Size, p.Error = ctx.URLParamInt("size")
		if p.Error != nil {
			return p
		}
	}
	if p.Size < 2 {
		p.Size = 2
	}
	if p.Size > p.Total {
		p.Size = p.Total
	}
	p.Pages = p.Total / p.Size
	if p.Total%p.Size > 0 {
		p.Pages += 1
	}
	if p.Page > p.Pages {
		p.Page = p.Pages
	}
	p.Start = (p.Page - 1) * p.Size
	p.End = p.Page * p.Size
	if p.End > p.Total {
		p.End = p.Total
	}
	return p
}
