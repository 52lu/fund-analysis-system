package main

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/initialize"
)

func init() {
	// 初始化全局配置文件
	initialize.InitConfig()
	// 初始化zap日志
	initialize.InitLogger()
	// 初始化gorm
	initialize.InitGorm()
	// 初始化redis
	initialize.InitRedis()
	// 初始化es
	initialize.InitES()
	// 定时任务
	initialize.InitCron()
}

func main() {
	// 程序退出前释放资源
	defer closeResource()
	// 启动服务
	RunServer()
}
// 程序退出前释放资源
func closeResource()  {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}
