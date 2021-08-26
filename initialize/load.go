package initialize
// 启动加载
func SetLoadInit() {
	// 初始化全局配置文件
	initConfig()
	// 初始化zap日志
	initLogger()
	// 初始化gorm
	initGorm()
	// 初始化redis
	initRedis()
	// 初始化es
	initElastic()
	// 定时任务
	initCron()
}
