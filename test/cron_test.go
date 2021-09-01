// Package test: 定时任务测试
package test

import (
	"52lu/fund-analye-system/crontab"
	"52lu/fund-analye-system/initialize"
	"testing"
)

func TestFundTopCron(t *testing.T) {
	initialize.SetLoadInit()
	topCron := crontab.FundTopCron{}
	topCron.Run()
}
