package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/log"
	"github.com/sirlanri/iot1-server/sqls"
)

//SendHumiTemp 接收树莓派发来的温湿度信息
func SendHumiTemp(con iris.Context) {
	humi := con.URLParam("humi")

	Humi = humi
	temp := con.URLParam("temp")

	Temp = temp
	Count++
	//如果满60次，写入数据库
	if Count == 6 {
		log.Log.Debugln("次数满 开始写入数据库")

		sqls.TempRes(temp)
		sqls.HumiRes(humi)

		Count = 0
	}
	//strTemp := strconv.FormatFloat(temp, 'f', 2, 64)
	//strHumi := strconv.FormatFloat(humi, 'f', 2, 64)
	words := fmt.Sprintf("温湿度接收成功\n温度：%.2s\n湿度%.2s", temp, humi)
	con.WriteString("Temp Humi confirmed")
	SendMqttInfo(words)

}

func Test(con iris.Context) {
	sqls.Mytest()
	con.WriteString("Success")
}

//SendTemps 接收树莓派发送温度*3
func SendTemps(con iris.Context) {
	var data TempStruct
	err := con.ReadJSON(&data)
	if err != nil {
		log.Log.Warnln("接收树莓派温度*3，传入数据错误", err.Error())
		con.StatusCode(iris.StatusBadRequest)
		con.WriteString("传入温度*3格式不正确")
		return
	}
	log.Log.Debugln("接收树莓派温度*3成功", data.Temp1, data.Temp2, data.Temp3)
	Temps = data
	con.WriteString("server已接收温度*3")

	sqls.TempRes(data.Temp1)
}

//接收单个温度数据
func SendSingleTemp(con iris.Context) {
	name := con.URLParam("name")
	num := con.URLParam("temp")
	if num == "" {
		con.WriteString("传入数据空")
		return
	}
	switch name {
	case "esp1":
		Temps.Temp1 = num
	case "esp2":
		Temps.Temp2 = num
	case "esp3":
		Temps.Temp3 = num
	}
	con.WriteString("Temp confirmed")
	sqls.TempRes(num)
}

//接收单个潮湿度数据
func SendSingleHumi(con iris.Context) {
	name := con.URLParam("name")
	num := con.URLParam("humi")
	if num == "" {
		con.WriteString("传入数据空")
		return
	}
	switch name {
	case "esp1":
		Humis.Humi1 = num
	case "esp2":
		Humis.Humi2 = num
	case "esp3":
		Humis.Humi3 = num
	}
	con.WriteString("Temp confirmed")
	sqls.HumiRes(num)
}

//SendHumis 接收树莓派发送的湿度*3
func SendHumis(con iris.Context) {
	var data HumiStruct
	err := con.ReadJSON(&data)
	if err != nil {
		log.Log.Warnln("接收树莓派湿度*3 传入错误", err.Error())
		con.StatusCode(iris.StatusBadRequest)
		con.WriteString("传入湿度*3格式不正确")
		return
	}
	log.Log.Debugf("接收树莓派潮湿度*3成功 %s %s %s", data.Humi1, data.Humi2, data.Humi3)
	Humis = data
	con.WriteString("server已接收潮湿度*3")
	sqls.HumiRes(data.Humi1)

}

// 接收树莓派发送的雨量&增量
func SendRain(con iris.Context) {
	rain, err := con.URLParamFloat64("rain")
	if err != nil {
		log.Log.Warnln("传入雨量数据错误")
		con.StatusCode(400)
		con.WriteString("传入雨量数据错误")
		return
	}

	RainInc = Rain - rain
	Rain = rain
	con.WriteString("已接收雨量&增量")
}

//接收树莓派发送的水量
func SendWater(con iris.Context) {
	water := con.URLParam("water")
	if water == "" {
		con.StatusCode(400)
		con.WriteString("传入水量数据错误")
		return
	}
	Water = water
	con.WriteString("接收水量成功")
}

//接收树莓派发送的光照强度
func SendLight(con iris.Context) {
	light := con.URLParam("light")
	if light == "" {
		log.Log.Warn("树莓派传入光照强度出错 ")
		con.StatusCode(400)
		con.WriteString("树莓派传入光照强度出错")
		return
	}
	Light = light
	con.WriteString("接收光照成功")
}
