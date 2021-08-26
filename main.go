package main

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/initialize"
	_ "52lu/fund-analye-system/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 程序退出前释放资源
	defer closeResource()
	// 加载启动前配置
	initialize.SetLoadInit()
	// 启动服务
	RunServer()
}

// 程序退出前释放资源
func closeResource() {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}
