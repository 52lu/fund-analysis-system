/**
 * @Author Mr.LiuQH
 * @Description 路由注册入口
 * @Date 2021/7/5 3:17 下午
 **/
package initialize

import (
	"52lu/fund-analye-system/router/demo"
	"github.com/gin-gonic/gin"
)

// 注册路由入口
func RegisterRouters(engine *gin.Engine) {
	// 注册演示路由
	demo.InitDemoRouter(engine)
}
