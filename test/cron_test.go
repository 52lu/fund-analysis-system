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
