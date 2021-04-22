package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-server/sqls"
)

//SendHumiTemp 接收树莓派发来的温湿度信息
func SendHumiTemp(con iris.Context) {
	humi, err := con.URLParamFloat64("humi")
	if err != nil {
		fmt.Println("sensor 传入湿度错误 ", err.Error())
		return
	}
	Humi = humi
	temp, err := con.URLParamFloat64("temp")
	if err != nil {
		fmt.Println("sensor 传入温度错误 ", err.Error())
		return
	}
	Temp = temp
	Count++
	//如果满60次，写入数据库
	if Count == 5 {
		fmt.Println("次数满 开始写入数据库")
		go func() {
			sqls.TempRes(float32(temp))
			sqls.HumiRes(float32(humi))
		}()
		Count = 0
	}

}
