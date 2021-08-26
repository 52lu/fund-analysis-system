package demo

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/internal"
	"52lu/fund-analye-system/internal/validate"
	user2 "52lu/fund-analye-system/model/entity/user"
	"52lu/fund-analye-system/model/request/user"
	"52lu/fund-analye-system/model/response"
	userService "52lu/fund-analye-system/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @description: 用户注册
 * @param ctx
 */
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam user.RegisterParam
	var err error
	_ = ctx.ShouldBindJSON(&registerParam)
	//  参数校验
	err = validate.Validate(registerParam)
	if err != nil {
		response.Error(ctx, "参数校验失败: " + err.Error())
		return
	}
	// 调用注册
	register, err := userService.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败: " + err.Error())
		return
	}
	response.OkWithData(ctx, register)
}

/**
 * @description: 用户账号密码登录
 * @param ctx
 */
func Login(ctx *gin.Context) {
	// 绑定参数
	var loginParam user.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	//  参数校验
	err := validate.Validate(loginParam)
	if err != nil {
		response.Error(ctx, "参数校验失败: " + err.Error())
		return
	}
	// 调用登录服务
	userRecord := user2.User{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := userService.LoginPwd(&userRecord); err != nil {
		global.GvaLogger.Error("登录失败:", zap.Any("user", userRecord))
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	// 生成token
	token, err := internal.CreateToken(userRecord.ID)
	if err != nil {
		global.GvaLogger.Sugar().Errorf("登录失败,Token生成异常:%s", err)
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	userRecord.Token = token
	response.OkWithData(ctx, userRecord)
}

// 查询用户信息
func GetUser(ctx *gin.Context) {
	// 从上下文中获取用户信息，(经过中间件逻辑时，已经设置到上下文)
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
	return
}
