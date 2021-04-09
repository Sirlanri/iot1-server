package sqls

import "fmt"

//BodyRes SQL 写入人体传感器的数据
func BodyRes(resFlag int) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into bodysensor (status) 
		values (?)`, resFlag)
	if err != nil {
		fmt.Println("人体传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("人体传感器，commit出错", err.Error())
		return false
	}
	return true
}

//TempRes -SQL 写入温度数据 float
func TempRes(temp float32) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into tempsensor (num)
		values (?)`, temp)
	if err != nil {
		fmt.Println("温度传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("温度传感器，commit出错", err.Error())
		return false
	}
	return true
}

//HumiRes -SQL 写入湿度数据 float
func HumiRes(humi float32) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into humisensor (num)
		values (?)`, humi)
	if err != nil {
		fmt.Println("湿度传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("湿度传感器，commit出错", err.Error())
		return false
	}
	return true
}

//LightRes -SQL 光照传感器 写入光照数据
func LightRes(light float32) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into lightsensor (num)
		values (?)`, light)
	if err != nil {
		fmt.Println("光照传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("光照传感器，commit出错", err.Error())
		return false
	}
	return true
}

// VoiceRes -SQL 声音传感器 写入
func VoiceRes(voice float64) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`insert into voicesensor (num)
		values (?)`, voice)
	if err != nil {
		fmt.Println("声音传感器，写入出错", err.Error())
		return false
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("声音传感器，commit出错", err.Error())
		return false
	}
	return true
}

//GetTimePer -SQL 获取有无人的次数
func GetTimePer() (have, no int) {
	tx, _ := Db.Begin()
	rows, err := tx.Query(`SELECT COUNT(*) FROM bodysensor
		WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=1 
		UNION 
		SELECT COUNT(*) FROM bodysensor
		WHERE itime>=DATE_SUB(now(),interval 1 day) AND status=0;`)
	if err != nil {
		fmt.Println("查询有无人出错", err.Error())
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
			fmt.Println(err.Error())
			panic("写入int出错")
		}
	}
	return
}

//GetWeekTempHumi 获取一周中每天的温度&湿度平均值
func GetWeekTempHumi() (data map[string][]float32) {
	tx, _ := Db.Begin()
	tempRows, err := tx.Query(`SELECT round(AVG(num),2) FROM tempsensor 
		WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY day(itime) ORDER BY day(itime);`)
	if err != nil {
		fmt.Println("查询温度平均值错误", err.Error())
	}
	var temps []float32
	for tempRows.Next() {
		var temp float32
		err = tempRows.Scan(&temp)
		if err != nil {
			fmt.Println("读取temp数据出错", err.Error())
		}
		temps = append(temps, temp)
	}

	humiRows, err := tx.Query(`SELECT round(AVG(num),2) FROM humisensor 
		WHERE itime>=DATE_SUB(now(),interval 7 day)
		GROUP BY day(itime) ORDER BY day(itime);`)
	if err != nil {
		fmt.Println("查询温度平均值错误", err.Error())
	}
	var humis []float32
	for humiRows.Next() {
		var humi float32
		err = humiRows.Scan(&humi)
		if err != nil {
			fmt.Println("读取humi数据出错", err.Error())
		}
		humis = append(humis, humi)
	}
	data = map[string][]float32{
		"temp": temps,
		"humi": humis,
	}
	return
}
