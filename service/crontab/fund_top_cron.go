package crontab

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/service/crawl/fund"
	"fmt"
)

type FundTopCron struct {}

func (c FundTopCron) Run()  {
	f := &fund.TopCrawl{}
	// 爬取网页
	f.CrawlHtml()
	// 转换数据
	fundDayTopList := f.ConvertEntity()
	// 入库
	result := global.GvaMysqlClient.Create(fundDayTopList)
	fmt.Println("保存错误:", result.Error)
	fmt.Println("保存结果:", result.RowsAffected)
}
