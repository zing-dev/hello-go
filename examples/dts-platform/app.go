package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"log"
	"sync"
)

type App struct {
	locker sync.Mutex
	dts    sync.Map
}

func (a *App) InitDTS() {
	d1 := NewDTS(DTSModel{Id: 1, ChannelId: 4, Name: "d1", Host: "192.168.0.86", SDKPort: 17083})
	d2 := NewDTS(DTSModel{Id: 2, ChannelId: 4, Name: "d2", Host: "192.168.0.215", SDKPort: 17083})
	d1.Run()
	d2.Run()
	a.dts.Store(d1.Model.Id, d1)
	a.dts.Store(d2.Model.Id, d2)
}

func (a *App) Close() {
	a.dts.Range(func(_, v interface{}) bool {
		v.(*DTS).Close()
		return true
	})
}

func (a *App) Run() {
	app := iris.New()

	a.InitDTS()

	app.Get("/", func(ctx context.Context) {
		_, _ = ctx.WriteString("success")
	})
	app.Get("/run", func(ctx context.Context) {
		id, err := ctx.URLParamInt("id")
		if err != nil {
			_, _ = ctx.WriteString("err:" + err.Error())
			return
		}
		if dts, ok := a.dts.Load(id); ok {
			dts.(*DTS).Run()
		}
		_, _ = ctx.WriteString("success run " + ctx.URLParam("id"))
	})
	app.Get("/close", func(ctx context.Context) {
		id, err := ctx.URLParamInt("id")
		if err != nil {
			_, _ = ctx.WriteString("err:" + err.Error())
			return
		}
		if dts, ok := a.dts.Load(id); ok {
			dts.(*DTS).Close()
		}
		_, _ = ctx.WriteString("success close " + ctx.URLParam("id"))
	})
	app.Get("/zones", func(ctx context.Context) {
		var list [][]Zone
		a.dts.Range(func(_, v interface{}) bool {
			dts := v.(*DTS)
			list = append(list, dts.Zones)
			return true
		})
		_, _ = ctx.JSON(list)
	})
	app.Get("/list", func(ctx context.Context) {
		var list []DTS
		a.dts.Range(func(_, v interface{}) bool {
			dts := v.(*DTS)
			list = append(list, DTS{DeviceId: dts.DeviceId, Status: dts.Status, Model: dts.Model})
			return true
		})
		_, _ = ctx.JSON(list)
	})
	err := app.Run(iris.Addr(":7777"))
	if err != nil {
		log.Fatal(err)
	}
}
