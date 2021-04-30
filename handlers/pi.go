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
		go func() {
			sqls.TempRes(temp)
			sqls.HumiRes(humi)
		}()
		Count = 0
	}
	//strTemp := strconv.FormatFloat(temp, 'f', 2, 64)
	//strHumi := strconv.FormatFloat(humi, 'f', 2, 64)
	words := fmt.Sprintf("温湿度接收成功\n温度：%.2f\n湿度%.2f", temp, humi)
	con.WriteString("Temp Humi confirmed")
	SendMqttInfo(words)

}

func Test(con iris.Context) {
	con.WriteString("Success")
}
