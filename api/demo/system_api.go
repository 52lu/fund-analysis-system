package demo

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/response"
	"github.com/gin-gonic/gin"
)

// 配置信息
func GetConfig(ctx *gin.Context)  {
	response.OkWithData(ctx,global.GvaConfig)
}