package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/log"
	"github.com/sirlanri/iot1-server/sqls"
)

//SendHumiTemp 接收树莓派发来的温湿度信息
func SendHumiTemp(con iris.Context) {
	humi, err := con.URLParamFloat64("humi")
	if err != nil {
		log.Log.Errorln("sensor 传入湿度错误 ", err.Error())
		return
	}
	Humi = humi
	temp, err := con.URLParamFloat64("temp")
	if err != nil {
		log.Log.Errorln("sensor 传入温度错误 ", err.Error())
		return
	}
	Temp = temp
	Count++
	//如果满60次，写入数据库
	if Count == 6 {
		log.Log.Debugln("次数满 开始写入数据库")
		go func() {
			res1 := sqls.TempRes(float32(temp))
			res2 := sqls.HumiRes(float32(humi))
			if res1 && res2 {
				log.Log.Debugln("数据库写入完毕")
			}
		}()
		Count = 0
	}
	//strTemp := strconv.FormatFloat(temp, 'f', 2, 64)
	//strHumi := strconv.FormatFloat(humi, 'f', 2, 64)
	words := fmt.Sprintf("温湿度接收成功\n温度：%.2f\n湿度%.2f", temp, humi)
	SendMqttInfo(words)

}
