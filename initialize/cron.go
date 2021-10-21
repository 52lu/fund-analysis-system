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
	_, _ = c.AddJob("@every 1h", crontab.DemoCron{})
	// 爬取每日基金排行榜单(每天 13:30)
	_, _ = c.AddJob("0 30 13 */1 * *", crontab.FundTopCron{})
	// 爬取基金基本信息(每天 1:30)
 	_, _ = c.AddJob("0 30 15 */1 * *", crontab.FundBasicCron{})
	// 爬取基金持仓信息信息(每天 20:30)
	_, _ = c.AddJob("0 30 20 */1 * *", crontab.FundStockCron{})

}

// 添加Func任务
func addFunc(c *cron.Cron) {

}
