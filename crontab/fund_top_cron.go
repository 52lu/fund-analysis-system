package crontab

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/service/crawl/fund"
	"fmt"
)

type FundTopCron struct {}

func (c FundTopCron) Run()  {
	fmt.Println("基金排行榜-定时任务准备运行....")
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
		global.GvaLogger.Sugar().Infof("本次任务保存数据：%条",result.RowsAffected)
		return
	}
	global.GvaLogger.Sugar().Info("任务运行成功，无数据要保存！")
	fmt.Println("基金排行榜-定时任务运行结束！")
}
