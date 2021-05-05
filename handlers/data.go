package handlers

var (
	Humi string //湿度data del
	Temp string //温度data

	Humis HumiStruct //记录湿度*3
	Temps TempStruct //记录温度*3

	Rain    string //雨量
	RainInc string //雨增量
	Count   int16  //计数
)

//温度*3
type TempStruct struct {
	Temp1 string `json:"temp1"`
	Temp2 string `json:"temp2"`
	Temp3 string `json:"temp3"`
}

//湿度*3
type HumiStruct struct {
	Humi1 string `json:"humi1"`
	Humi2 string `json:"humi2"`
	Humi3 string `json:"humi3"`
}
