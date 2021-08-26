/**
 * @Author Mr.LiuQH
 * @Description 响应统一封装
 * @Date 2021/7/5 4:12 下午
 **/
package response

import (
	"52lu/fund-analye-system/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	SUCCESS = 0
	ERROR   = -1
	TOKEN_EXIPRE   = -1
)

// 定义统一返回接口格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"-"`
}

// 请求响应
func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	// 格式化时间
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format(global.YYYYMMDDHHIISS),
	})
}

// 返回固定成功信息
func Ok(ctx *gin.Context) {
	ResultJson(ctx, SUCCESS, "请求成功", map[string]interface{}{})
}

// 只返回成功消息
func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCESS, msg, map[string]interface{}{})
}

// 返回固定消息和数据
func OkWithData(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, SUCCESS, "请求成功", data)
}

// 返回指定消息和数据
func OkWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ResultJson(ctx, SUCCESS, msg, data)
}

// 错误信息
func Error(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, map[string]interface{}{})
}

// 登录超时或者token失效
func ErrorWithToken(ctx *gin.Context, msg string)  {
	ResultJson(ctx, TOKEN_EXIPRE, msg, map[string]interface{}{})
}