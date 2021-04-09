package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/configs"
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
	front.Get("/getLight", handlers.GetLight)
	//front.Get("/getVoice", handlers.GetVoiceTest)
	front.Get("/getWeek", handlers.GetWeekData)
	front.Get("/getTimePer", handlers.GetTimePer)
	front.Get("/getWeekTempHumi", handlers.GetWeekTempHumi)

	//接收传感器发来的内容
	sensor := m2m.Party("/sensor")
	sensor.Post("/temp", handlers.SendTemp)
	sensor.Post("/light", handlers.SendLight)
	sensor.Post("/humi", handlers.SendHumi)
	sensor.Post("/body", handlers.SendBody)

	//向传感器发送数据
	sensor.Get("/lighton", handlers.Lighton)
	sensor.Get("/lightoff", handlers.Lightoff)
	sensor.Get("/buzzon", handlers.Buzzon)

	app.Run(iris.Addr(configs.PortConfig()))

	return
}
