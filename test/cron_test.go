// Package test: 定时任务测试
package test

import (
	"52lu/fund-analye-system/crontab"
	"testing"
)

func TestFundTopCron(t *testing.T) {
	topCron := crontab.FundTopCron{}
	topCron.Run()
}

// 批量抓取基金详情
func TestBatchCrawlDetailCron(t *testing.T) {
	cron := crontab.FundBasicCron{}
	cron.Run()
}

// 批量抓取基金持仓
func TestFundStockCron(t *testing.T) {
	cron := crontab.FundStockCron{}
	cron.Run()
}