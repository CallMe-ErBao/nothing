package config

import (
	"pandora/modules/db"
	"pandora/vars"
)

func Init() {
	mysqlUrl := vars.Conf.GetString("db")
	db.MysqlURL = mysqlUrl

}
