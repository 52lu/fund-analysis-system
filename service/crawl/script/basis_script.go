// crawl: 基金基础信息
package script

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

// 定义结构体对应
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
 * @description: 获取基金基本信息
 * @param fundCode
 */
func (f *FundBaseCrawl) GetFundBasis(fundCode string) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		global.GvaLogger.Sugar().Errorf("基金%s,信息获取失败: %s", fundCode, err)
		return
	})
	// 基金概况
	collector.OnHTML("div[class='txt_cont']", func(element *colly.HTMLElement) {
		err := element.Unmarshal(f)
		if err != nil {
			fmt.Println("element.Unmarshal error: ", err)
		}
		// 格式化数据
		f.format()
	})
	err := collector.Visit(fmt.Sprintf("https://fundf10.eastmoney.com/jbgk_%s.html", fundCode))
	if err != nil {
		global.GvaLogger.Sugar().Errorf("基金%s,信息请求失败: %s", fundCode, err)
	}
}

// 格式化数据
func (f *FundBaseCrawl) format() {
	f.Code = utils.ExtractNumberFromString(f.Code)
	// 发布时间
	f.ReleaseDate = replaceDateChinese(f.ReleaseDate)
	// 成立日期
	f.EstablishDate = strings.TrimSpace(replaceDateChinese(strings.Split(f.EstablishDate, "/")[0]))
	// 成立规模
	f.EstablishShares = utils.ExtractNumberFromString(replaceDateChinese(strings.Split(f.EstablishShares, "/")[1]))
	// 管理费率
	f.ManageFeeRate = utils.ExtractNumberFromString(f.ManageFeeRate)
	// 托管费率
	f.CustodyFeeRate = utils.ExtractNumberFromString(f.CustodyFeeRate)
	// 销售服务费率
	f.SaleFeeRate = utils.ExtractNumberFromString(f.SaleFeeRate)
}

// 处理日期，把年、月、日替换成-
func replaceDateChinese(strDate string) string {
	strDate = strings.ReplaceAll(strDate, "年", "-")
	strDate = strings.ReplaceAll(strDate, "月", "-")
	strDate = strings.ReplaceAll(strDate, "日", "")
	return strDate
}
