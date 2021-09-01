package initialize

import (
	"52lu/fund-analye-system/crontab"
	"52lu/fund-analye-system/global"
	"github.com/robfig/cron/v3"
)

// 定时任务管理
func initCron() {
	if !global.GvaConfig.Cron.Enable {
		return
	}
	c := cron.New(cron.WithSeconds())
	addJob(c)
	addFunc(c)
	c.Start()

}

// 添加Job任务
func addJob(c *cron.Cron) {
	// 测试任务
	//_, _ = c.AddJob("@every 1s", crontab.DemoCron{})
	// 爬取每日基金排行榜单
	_, _ = c.AddJob("0 30 16 */1 * *", crontab.FundTopCron{})
	//为了测试，先写成 每10s一次
 	//_, _ = c.AddJob("@every 10s", crontab.FundTopCron{})

}

// 添加Func任务
func addFunc(c *cron.Cron) {

}
