package handlers

import (
	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/sqls"
)

//GetWeekTempHumi 获取一周温度湿度的平均值列表
func GetWeekTempHumi(con iris.Context) {
	data := sqls.GetWeekTempHumi()
	con.JSON(data)
}

//GetRealtime 获取内存中的实时温湿度数据 del
func GetRealtime(con iris.Context) {
	data := map[string]string{
		"temp": Temp,
		"humi": Humi,
	}
	con.JSON(data)
	SendMqttInfo("获取实时数据")
}

//GetRealTemp 获取实时温度*3
func GetRealTemp(con iris.Context) {
	con.JSON(Temps)
}

//获取实时潮湿度*3
func GetRealHumi(con iris.Context) {
	con.JSON(Humis)
}

//获取实时雨量&增量
func GetRealRainInc(con iris.Context) {
	data := map[string]float64{
		"rain": Rain,
		"inc":  RainInc,
	}
	con.JSON(data)
}

//获取实时水量
func GetRealWater(con iris.Context) {
	con.WriteString(Water)
}

//Setled 前端控制LED
func Setled(con iris.Context) {
	auth := con.URLParam("auth")
	if auth != "iris" {
		con.StatusCode(201)
		con.WriteString("无权限操作")
		return
	}
	//操作码：on off blink
	code := con.URLParam("code")
	if code == "on" || code == "off" || code == "blink" {
		SendMqttIns(code, "pi/res/led")
		return
	}
	con.WriteString("传入数据不合法")

}
