package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/sqls"
)

//GetWeekTempHumi 获取一周温度湿度的平均值列表
func GetWeekTempHumi(con iris.Context) {
	data := sqls.GetWeekTempHumi()
	con.JSON(data)
}

//GetRealtime 获取内存中的实时温湿度数据
func GetRealtime(con iris.Context) {
	temp := fmt.Sprintf("%.2f", Temp)
	humi := fmt.Sprintf("%.2f", Humi)
	data := map[string]string{
		"temp": temp,
		"humi": humi,
	}
	con.JSON(data)
	SendMqttInfo("获取实时数据")
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
