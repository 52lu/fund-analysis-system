package crontab

import "52lu/fund-analye-system/service/crawl/script"

type FundCrawlCron struct {
	Code string
}

func (c FundCrawlCron) Run()  {
	// 爬取页面信息
	var fundBaseCrawl script.FundBaseCrawl
	fundBaseCrawl.GetFundBasis(c.Code)
	//
}