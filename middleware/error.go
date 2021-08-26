package middleware

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 捕获请求全局错误
func CatchErrorMiddleWare() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 捕获错误
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("运行异常: %s", err)
				// 捕获错误
				if global.GvaLogger != nil {
					global.GvaLogger.Error(errMsg)
				}
				// todo 邮件通知

				// 错误响应
				response.Error(ctx, errMsg)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
