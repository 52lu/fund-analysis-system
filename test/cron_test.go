// Package test: 定时任务测试
package test

import (
	"52lu/fund-analye-system/initialize"
	"52lu/fund-analye-system/service/crontab"
	"testing"
)

func TestFundTopCron(t *testing.T) {
	initialize.SetLoadInit()
	topCron := crontab.FundTopCron{}
	topCron.Run()
}
