package demo

import (
	"52lu/fund-analye-system/model/response"
	"52lu/fund-analye-system/service/crawl/fund"
	"github.com/gin-gonic/gin"
)
/**
*  Run
*  @Description:
*  @param context
**/
func Run(context *gin.Context){
	f := &fund.BasisCrawl{}
	f.CrawlHtml("001751")
	fundEntity := f.ConvertToEntity()
	response.OkWithData(context,fundEntity)
}
