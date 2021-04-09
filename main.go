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
	m2m := app.Party("/m2m", crs).AllowMethods(iris.MethodOptions)

	//前端
	front := m2m.Party("/front")
	//前端获取实时数据
	front.Get("/getTemp", handlers.GetTemp)
	front.Get("/getHumi", handlers.GetHumi)

	//接收传感器发来的内容
	sensor := m2m.Party("/sensor")
	sensor.Post("/humitemp", handlers.SendHumiTemp)

	//向传感器发送数据

	app.Run(iris.Addr(configs.PortConfig()))

	return
}
