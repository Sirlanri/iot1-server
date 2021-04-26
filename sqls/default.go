/*Package sqls 执行数据库的增删改查操作
 */
package sqls

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirlanri/iot1-server/configs"
	"github.com/sirlanri/iot1-server/log"
)

//Db 创建的唯一指针
var Db *sql.DB

//初始化，自动创建db指针
func init() {
	Db = ConnectDB()
}

//ConnectDB 初始化时，连接数据库
func ConnectDB() *sql.DB {
	database := configs.SQLConfg()
	Db, err := sql.Open("mysql", database)
	if err != nil {
		log.Log.Errorln("数据库初始化链接失败", err.Error())
	}

	if Db.Ping() != nil {
		log.Log.Errorln("初始化-数据库-用户/密码/库验证失败", Db.Ping().Error())
		return nil
	}

	return Db
}
