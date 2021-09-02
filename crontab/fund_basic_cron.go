package crontab

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/service/crawl/fund"
	"fmt"
	"time"
)

type FundBasicCron struct {
	Code string
}
// 抓取详情信息
func (c FundBasicCron) Run() {
	begin := time.Now().UnixMilli()
	fmt.Println("基金详情-定时任务开始运行")
	// 批量爬取
	fundBasicEntityList := fund.BatchBasicCrawl()
	if fundBasicEntityList != nil {
		// 保存入库
		create := global.GvaMysqlClient.Create(fundBasicEntityList)
		if create.Error != nil {
			global.GvaLogger.Sugar().Errorf("基金详情入库失败",create.Error)
			return
		}
		global.GvaLogger.Sugar().Infof("基金详情抓取成功，共: %v 条",create.RowsAffected)
	}
	fmt.Printf("基金详情-定时任务运行完成,耗时:%vms\n",time.Now().UnixMilli() - begin)
}
