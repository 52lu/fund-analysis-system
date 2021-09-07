// Package fund: 基金股票持仓
package fund

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/service/crawl"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 对应表中的每一tr
type StockPercentageRow struct {
	StockCode  string `selector:"td:nth-of-type(2)"`
	StockHref  string `selector:"td:nth-of-type(2) > a[href]" attr:"href"`
	StockName  string `selector:"td:nth-of-type(3)"`
	Percentage string `selector:"td:nth-of-type(7)"`
	Quantity   string `selector:"td:nth-of-type(8)"`
	Amount     string `selector:"td:nth-of-type(9)"`
}

// 对应整个table
type StockPercentageRowsCrawl struct {
	Rows       []StockPercentageRow `selector:"tr"`
	FundCode   string
	CutOffDate string
}

// 爬取信息
func (c *StockPercentageRowsCrawl) CrawlHtml(fundCode string) {
	collector := colly.NewCollector(colly.UserAgent(crawl.UserAgent), colly.Async(true))
	// 开启限速
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*fundf10.eastmoney.*",
		Delay:       500 * time.Millisecond,
		RandomDelay: 500 * time.Millisecond,
		Parallelism: 20,
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("url:", request.URL)
	})
	// 处理返回的数据
	collector.OnResponse(func(response *colly.Response) {
		// 替换字符串
		compile := regexp.MustCompile(`var apidata=\{ content:"(.*)",arryear:`)
		matchResult := compile.FindAllStringSubmatch(string(response.Body), -1)
		if len(matchResult) == 0 {
			return
		}
		htmlString := matchResult[0][1]
		htmlString = strings.ReplaceAll(htmlString, "%", "")
		htmlString = strings.ReplaceAll(htmlString, ",", "")
		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(htmlString)))
		if err != nil {
			return
		}
		docSelection := doc.Find("div[class='box']").First()
		e := &colly.HTMLElement{
			DOM: docSelection.Find("table"),
		}
		err = e.Unmarshal(c)
		if err != nil {
			global.GvaLogger.Error("爬虫解析失败", zap.String("error", err.Error()))
			return
		}
		// 过滤header
		if len(c.Rows) > 0 && c.Rows[0].StockCode == "" {
			c.Rows = c.Rows[1:]
		}
		// 获取持仓季度时间信息
		c.CutOffDate = docSelection.Find("h4 label").Eq(1).Find("font").Text()
		// 补充额外信息
		c.FundCode = fundCode
	})
	err = collector.Visit(fmt.Sprintf("https://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=%s&topline=30", fundCode))
	if err != nil {
		global.GvaLogger.Sugar().Errorf("CrawlHtml error:%s", err)
	}
	collector.Wait()
}

// 数据清洗
func (c StockPercentageRowsCrawl) ConvertEntity() []entity.FundStock {
	var fundStocks []entity.FundStock
	if len(c.Rows) < 1 {
		return []entity.FundStock{}
	}
	for _, row := range c.Rows {
		item := entity.FundStock{
			FundCode:   c.FundCode,
			StockCode:  row.StockCode,
			StockName:  row.StockName,
			CutOffDate: c.CutOffDate,
		}
		// 提取交易所信息
		// 提取交易所信息
		compile := regexp.MustCompile(`com\/([a-zA-Z]+)\d+\.html`)
		stringSubMatch := compile.FindAllStringSubmatch(row.StockHref, -1)
		if stringSubMatch != nil {
			 item.StockExchange = strings.ToUpper(stringSubMatch[0][1])
		}
		// 字符串转浮点型
		item.Percentage, _ = strconv.ParseFloat(row.Percentage, 64)
		item.Quantity, _ = strconv.ParseFloat(row.Quantity, 64)
		item.Amount, _ = strconv.ParseFloat(row.Amount, 64)
		fundStocks = append(fundStocks, item)
	}
	return fundStocks
}
