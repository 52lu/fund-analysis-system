// test: 临时测试
package test

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/initialize"
	"52lu/fund-analye-system/utils"
	"fmt"
	"testing"
)


func TestRun(t *testing.T) {
	existDatabase()
}

func existDatabase()  {
	initialize.SetLoadInit()
	var  databases []string
	global.GvaMysqlClient.Raw("show DATABASES").Scan(&databases)
	dbName := global.GvaConfig.Mysql.Database
	if _, ok := utils.ExistSliceStr(dbName, databases); !ok {
		sql := fmt.Sprintf("create database `%s` character set utf8mb4 collate utf8mb4_bin",dbName)
		global.GvaMysqlClient.Exec(sql)
	}
}
