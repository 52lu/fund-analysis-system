package crontab

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/service/crawl/fund"
	"fmt"
)

type FundCrawlCron struct {
	Code string
}

func (c FundCrawlCron) Run() {
	f := &fund.BaseCrawl{}
	// 爬取页面信息
	f.CrawlHtml(c.Code)
	// 转成实体类型
	fundBasicEntity := &entity.FundBasis{}
	f.ConvertToEntity(fundBasicEntity)
	fmt.Println("GvaMysqlClient",global.GvaMysqlClient)
	// 保存入库
	create := global.GvaMysqlClient.Create(fundBasicEntity)
	fmt.Println("保存结果:", create.Error)
	fmt.Println("保存结果:", create.RowsAffected)

}
