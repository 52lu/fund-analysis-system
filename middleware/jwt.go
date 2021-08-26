package middleware

import (
	"52lu/fund-analye-system/internal"
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/dao"
	"52lu/fund-analye-system/model/request/user"
	"52lu/fund-analye-system/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @description: JWT中间件
 * @return func(ctx *gin.Context)
 */
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取参数中的token
		token := getToken(ctx)
		global.GvaLogger.Sugar().Infof("token: %s", token)
		if token == "" {
			response.Error(ctx, "Token不能为空!")
			// 中断请求
			ctx.Abort()
			return
		}
		// 验证Token
		userClaim, err := internal.ParseToken(token)
		if err != nil {
			response.ErrorWithToken(ctx, "Token error :"+err.Error())
			// 中断请求
			ctx.Abort()
			return
		}
		// 设置到上下文中
		setContextData(ctx, userClaim, token)
		// 继续请求后续流程
		ctx.Next()
	}
}
// 设置数据到上下文
func setContextData(ctx *gin.Context, userClaim *user.UserClaims, token string) {
	userDao := &dao.UserDao{
		Uid: userClaim.Uid,
	}
	user, err := userDao.FindUser()
	if err != nil {
		response.Error(ctx, "用户不存在!")
		// 中断请求
		ctx.Abort()
		return
	}
	user.Token = token
	ctx.Set("userClaim", userClaim)
	ctx.Set("user", user)
}

// 从请求中获取Token
func getToken(ctx *gin.Context) string {
	var token string
	// 从header中获取
	token = ctx.Request.Header.Get("TOKEN")
	if token != "" {
		return token
	}
	// 获取当前请求方法
	if ctx.Request.Method == http.MethodGet {
		// 从Get请求中获取Token
		token, ok := ctx.GetQuery("token")
		if ok {
			return token
		}
	}
	// 从POST中和获取
	if ctx.Request.Method == http.MethodPost {
		// 从Get请求中获取Token
		postParam := make(map[string]interface{})
		_ = ctx.ShouldBindJSON(&postParam)
		token, ok := postParam["token"]
		if ok {
			return token.(string)
		}
	}
	return ""
}
