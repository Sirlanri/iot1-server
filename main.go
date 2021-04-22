package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/configs"
	"github.com/sirlanri/iot1-server/handlers"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, handlers.NotFound)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	iot1 := app.Party("/iot1", crs).AllowMethods(iris.MethodOptions)

	//前端
	front := iot1.Party("/web")
	//前端获取实时数据
	front.Get("/getRealtime", handlers.GetRealtime)
	front.Get("/setled", handlers.Setled)
	//接收传感器发来的内容
	sensor := iot1.Party("/sensor")
	sensor.Post("/humitemp", handlers.SendHumiTemp)

	//向传感器发送数据

	app.Run(iris.Addr(configs.PortConfig()))

	return
}
