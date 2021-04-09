package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/sqls"
)

//GetTimePer 数据统计页面，获取有无人的统计次数
func GetTimePer(con iris.Context) {
	have, no := sqls.GetTimePer()
	data := map[string]int{
		"have": have,
		"no":   no,
	}
	con.JSON(data)
}

//GetWeekTempHumi 获取一周温度湿度的平均值列表
func GetWeekTempHumi(con iris.Context) {
	data := sqls.GetWeekTempHumi()
	con.JSON(data)
}

//GetTemp 获取当前温度
func GetTemp(con iris.Context) {
	temp := fmt.Sprintf("%.2f", Temp)
	data := map[string]string{
		"temp": temp,
	}
	con.JSON(data)
}

//GetHumi 获取当前湿度
func GetHumi(con iris.Context) {
	humi := fmt.Sprintf("%.2f", Humi)
	data := map[string]string{
		"humi": humi,
	}
	con.JSON(data)
}

//GetLight 获取当前光强
func GetLight(con iris.Context) {
	light := Light
	data := map[string]int{
		"light": light,
	}
	con.JSON(data)
}
