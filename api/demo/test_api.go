package demo

import (
	"52lu/fund-analye-system/crontab"
	"52lu/fund-analye-system/model/response"
	"52lu/fund-analye-system/service/crawl/fund"
	"github.com/gin-gonic/gin"
)
/**
*  Run
*  @Description:
*  @param context
**/
func Run(ctx *gin.Context){
	code, b := ctx.GetQuery("code")
	if !b {
		response.Error(ctx,"参数不能为空!")
		return
	}
	f := &fund.BasisCrawl{}
	f.CrawlHtml(code)
	fundEntity := f.ConvertToEntity()
	response.OkWithData(ctx,fundEntity)
}

func Cron(ctx *gin.Context)  {
	query, _ := ctx.GetQuery("type")
	if query == "1" {
		fund.BatchBasicCrawl()
	} else {
		c := new(crontab.FundStockCron)
		c.Run()
	}
	response.Ok(ctx)
}