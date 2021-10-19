package crontab

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/service/crawl/fund"
)

type FundTopCron struct {}

func (c FundTopCron) Run()  {
	f := &fund.TopCrawlService{}
	// 爬取网页
	f.CrawlHtml()
	// 转换数据
	fundDayTopList := f.ConvertEntity()
	// 入库
	if !f.ExistTopDate() {
		result := global.GvaMysqlClient.Create(fundDayTopList)
		if result.Error != nil {
			global.GvaLogger.Sugar().Errorf("本次任务保存数据失败：%条",result.Error)
			return
		}
		global.GvaLogger.Sugar().Infof("本次任务保存数据：%d条",result.RowsAffected)
	}
}
