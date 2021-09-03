package crontab

import (
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
	// 开始爬取
	fund.BatchBasicCrawl()
	fmt.Printf("基金详情-定时任务运行完成,耗时:%vms\n",time.Now().UnixMilli() - begin)
}
