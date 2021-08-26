package main

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 获取自定义HTTP SERVER
func getCustomHttpServer(engine *gin.Engine) *http.Server {
	// 创建自定义配置服务
	httpServer := &http.Server{
		//ip和端口号
		Addr: global.GvaConfig.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		//计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		//计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		//请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

// RunServer 启动服务
func RunServer() {
	engine := gin.New()
	// 注册公共中间件
	engine.Use(gin.Recovery(), middleware.CatchErrorMiddleWare())
	// 获取自定义http配置
	httpServer := getCustomHttpServer(engine)
	// 注册路由
	registerRouters(engine)
	// 打印服务信息
	printServerInfo()
	// 启动服务
	err := httpServer.ListenAndServe()
	if err != nil {
		panic("启动失败: " + err.Error())
	}
}

// 打印服务信息
func printServerInfo() {
	appConfig := global.GvaConfig.App
	fmt.Printf("\n【 当前环境: %s 当前版本: %s 接口地址: http://%s 】\n", appConfig.Env, appConfig.Version, appConfig.Addr)
}
