package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

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
}
