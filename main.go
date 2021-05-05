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
	front.Get("/getRealTemp", handlers.GetRealTemp)
	front.Get("/getRealHumi", handlers.GetRealHumi)
	front.Get("/getRealRainInc", handlers.GetRealRainInc)
	front.Get("/setled", handlers.Setled)
	front.Get("/getWeekdata", handlers.GetWeekTempHumi)

	//接收传感器发来的内容
	sensor := iot1.Party("/sensor")
	sensor.Get("/humitemp", handlers.SendHumiTemp)
	sensor.Get("/test", handlers.Test)

	//接收树莓派的数据
	pi := iot1.Party("/pi")
	pi.Post("/temps", handlers.SendTemps)
	pi.Post("/humis", handlers.SendHumis)
	pi.Get("/rain", handlers.SendRain)

	app.Run(iris.Addr(configs.PortConfig()))

	return
}
