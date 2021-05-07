package sqls

import "github.com/sirlanri/iot1-server/log"

//TempRes -SQL 写入温度数据 float
func TempRes(temp string) bool {
	log.Log.Debug("开始写入温度")
	_, err := Db.Exec(`insert into tempsensor (num)
	values (?)`, temp)
	if err != nil {
		log.Log.Errorln("Temp写入数据库 初始化出错", err.Error())
		return false
	}

	log.Log.Debugln("Temp SQL写入完成")
	return true
}

//HumiRes -SQL 写入湿度数据 float
func HumiRes(humi string) bool {

	_, err := Db.Exec(`insert into humisensor (num)
		values (?)`, humi)
	if err != nil {
		log.Log.Errorln("湿度传感器，写入出错", err.Error())
		return false
	}
	log.Log.Debugln("Humi SQL写入完成")
	return true
}

//LightRes -SQL 光照传感器 写入光照数据
func LightRes(light float32) bool {
	_, err := Db.Exec(`insert into lightsensor (num)
		values (?)`, light)
	if err != nil {
		log.Log.Errorln("光照传感器，写入出错", err.Error())
		return false
	}

	return true
}

//GetTimePer -SQL 获取有无人的次数
func GetTimePer() (have, no int) {
	rows, err := Db.Query(`SELECT COUNT(*) FROM bodysensor
		WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=1 
		UNION 
		SELECT COUNT(*) FROM bodysensor
		WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=0;`)
	if err != nil {
		log.Log.Errorln("查询有无人出错", err.Error())
		return 0, 0
	}
	flag := true
	for rows.Next() {
		if flag {
			err = rows.Scan(&have)
			flag = false
		} else {
			err = rows.Scan(&no)
		}
		if err != nil {
			log.Log.Errorln(err.Error())
			panic("写入int出错")
		}
	}
	return
}

//GetWeekTempHumi 获取一周中每天的温度&湿度平均值
func GetWeekTempHumi() (data map[string][]float32) {
	tempRows, err := Db.Query(`SELECT round(AVG(num),2) FROM tempsensor 
		WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY day(itime) ORDER BY day(itime);`)
	defer tempRows.Close()
	if err != nil {
		log.Log.Errorln("查询温度平均值错误", err.Error())
	}
	var temps []float32
	for tempRows.Next() {
		var temp float32
		err = tempRows.Scan(&temp)
		if err != nil {
			log.Log.Errorln("读取temp数据出错", err.Error())
		}
		temps = append(temps, temp)
	}

	humiRows, err := Db.Query(`SELECT round(AVG(num),2) FROM humisensor 
		WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY day(itime) ORDER BY day(itime);`)
	defer humiRows.Close()
	if err != nil {
		log.Log.Errorln("查询温度平均值错误", err.Error())
	}
	var humis []float32
	for humiRows.Next() {
		var humi float32
		err = humiRows.Scan(&humi)
		if err != nil {
			log.Log.Errorln("读取humi数据出错", err.Error())
		}
		humis = append(humis, humi)
	}
	data = map[string][]float32{
		"temp": temps,
		"humi": humis,
	}
	return
}
