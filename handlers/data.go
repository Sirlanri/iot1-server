package handlers

var (
	Humi  string //湿度data
	Temp  string //温度data
	Count int16  //计数
)

//温度*3
type Temps struct {
	Temp1 string `json:"temp1"`
	Temp2 string `json:"temp2"`
	Temp3 string `json:"temp3"`
}

//湿度*3
type Humis struct {
	Humi1 string `json:"humi1"`
	Humi2 string `json:"humi2"`
	Humi3 string `json:"humi3"`
}
