// crawl: 基金基础信息
package crawl

import (
	"52lu/fund-analye-system/global"
	"fmt"
	"github.com/gocolly/colly/v2"
)

type FundBaseCrawl struct {
	Code            string `selector:"tr:nth-child(2) > td:nth-of-type(1)"`
	FullName        string `selector:"tr:nth-child(1) > td:nth-of-type(1)"`
	ShortName       string `selector:"tr:nth-child(1) > td:nth-of-type(2)"`
	Type            string `selector:"tr:nth-child(2) > td:nth-of-type(2)"`
	ReleaseDate     string `selector:"tr:nth-child(3) > td:nth-of-type(1)"`
	EstablishDate   string `selector:"tr:nth-child(3) > td:nth-of-type(2)"`
	EstablishShares string `selector:"tr:nth-child(3) > td:nth-of-type(2)"`
	Company         string `selector:"tr:nth-child(5) > td:nth-of-type(1)"`
	ManageFeeRate   string `selector:"tr:nth-child(7) > td:nth-of-type(1)"`
	CustodyFeeRate  string `selector:"tr:nth-child(7) > td:nth-of-type(2)"`
	SaleFeeRate     string `selector:"tr:nth-child(8) > td:nth-of-type(1)"`
	Benchmark       string `selector:"tr:nth-child(10) > td:nth-of-type(1)"`
}

/**
 * @description: 获取基金信息
 * @param fundCode
 */
func GetFundBase(fundCode string) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		global.GvaLogger.Sugar().Errorf("基金%s,信息获取失败: %s", fundCode, err)
		return
	})
	var fundBase FundBaseCrawl
	// 基金概况
	collector.OnHTML("div[class='txt_cont']", func(element *colly.HTMLElement) {
		err := element.Unmarshal(&fundBase)
		if err != nil {
			fmt.Println("element.Unmarshal error: ", err)
		}
		fmt.Printf("fundBase:%+v\n", fundBase)
	})
	// 基金持仓

	err := collector.Visit(fmt.Sprintf("https://fundf10.eastmoney.com/jbgk_%s.html", fundCode))
	if err != nil {
		global.GvaLogger.Sugar().Errorf("基金%s,信息请求失败: %s", fundCode, err)
	}
}