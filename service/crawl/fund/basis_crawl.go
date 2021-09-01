package fund

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
)

// BaseCrawl  定义结构体对应
type BaseCrawl struct {
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

// CrawlHtml 抓取取基金基本信息
func (f *BaseCrawl) CrawlHtml(fundCode string) {
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
	})
	err := collector.Visit(fmt.Sprintf("https://fundf10.eastmoney.com/jbgk_%s.html", fundCode))
	if err != nil {
		global.GvaLogger.Sugar().Errorf("基金%s,信息请求失败: %s", fundCode, err)
	}
}

// ConvertToEntity 格式化数据为实体类
func (f *BaseCrawl) ConvertToEntity(fundBaseEntity *entity.FundBasis) {
	fundBaseEntity.Code = utils.ExtractNumberFromString(f.Code)
	fundBaseEntity.FullName = f.FullName
	fundBaseEntity.ShortName = f.ShortName
	fundBaseEntity.Type = f.Type
	fundBaseEntity.Company = f.Company
	fundBaseEntity.Benchmark = f.Benchmark
	// 发布时间
	fundBaseEntity.ReleaseDate = replaceDateChinese(f.ReleaseDate)
	// 成立日期
	fundBaseEntity.EstablishDate = strings.TrimSpace(replaceDateChinese(strings.Split(f.EstablishDate, "/")[0]))
	// 成立规模
	establishShares := utils.ExtractNumberFromString(replaceDateChinese(strings.Split(f.EstablishShares, "/")[1]))
	fundBaseEntity.EstablishShares, _ = strconv.ParseFloat(establishShares, 64)
	// 管理费率
	manageFeeRate := utils.ExtractNumberFromString(f.ManageFeeRate)
	fundBaseEntity.ManageFeeRate, _ = strconv.ParseFloat(manageFeeRate, 64)
	// 托管费率
	fundBaseEntity.CustodyFeeRate, _ = strconv.ParseFloat(utils.ExtractNumberFromString(f.CustodyFeeRate), 64)
	// 销售服务费率
	fundBaseEntity.SaleFeeRate, _ = strconv.ParseFloat(utils.ExtractNumberFromString(f.SaleFeeRate), 64)
}

// 处理日期，把年、月、日替换成-
func replaceDateChinese(strDate string) string {
	strDate = strings.ReplaceAll(strDate, "年", "-")
	strDate = strings.ReplaceAll(strDate, "月", "-")
	strDate = strings.ReplaceAll(strDate, "日", "")
	return strDate
}


