package fund

import (
	"52lu/fund-analye-system/global"
	fundDao "52lu/fund-analye-system/model/dao/fund"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/service/crawl"
	"52lu/fund-analye-system/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type fundItem struct {
	FundCode         string `selector:"td:nth-of-type(1)"`
	FundName         string `selector:"td:nth-of-type(2)"`
	NetWorth         string `selector:"td:nth-of-type(3) > span.fb"`
	TopDate          string `selector:"td:nth-of-type(3) > span.date"`
	DayChange        string `selector:"td:nth-of-type(4)"`
	WeekChange       string `selector:"td:nth-of-type(5)"`
	MouthChange      string `selector:"td:nth-of-type(6)"`
	ThreeMouthChange string `selector:"td:nth-of-type(7)"`
	SixMouthChange   string `selector:"td:nth-of-type(8)"`
	YearChange       string `selector:"td:nth-of-type(9)"`
	TwoYearChange    string `selector:"td:nth-of-type(10)"`
	ThreeYearChange  string `selector:"td:nth-of-type(11)"`
	CurrentChange    string `selector:"td:nth-of-type(12)"`
	CreateChange     string `selector:"td:nth-of-type(13)"`
}

type TopCrawlService struct {
	Item []*fundItem `selector:"tr"`
}

// 检查该日期是否已经入库
func (f *TopCrawlService) ExistTopDate() bool {
	if len(f.Item) == 0 {
		return false
	}
	format := time.Now().Format("2006")
	topDate := fmt.Sprintf("%s-%s", format, f.Item[1].TopDate)
	fTopEntity, err := fundDao.FindLastOneByDate(topDate)
	if err != nil {
		global.GvaLogger.Error("查询数据库异常", zap.String("error", err.Error()))
	}
	if fTopEntity.ID != 0 {
		return true
	}
	return false
}

// CrawlHtml 抓取取基金基本信息
func (f *TopCrawlService) CrawlHtml() {
	collector := colly.NewCollector(
		colly.UserAgent(crawl.UserAgent),
	)
	// 设置Header
	collector.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")

	})
	collector.OnError(func(response *colly.Response, err error) {
		global.GvaLogger.Sugar().Errorf("基金排行榜,信息获取失败: %s", err)
		return
	})
	collector.OnHTML("#tblite_hh", func(element *colly.HTMLElement) {
		err := element.Unmarshal(f)
		if err != nil {
			fmt.Println("element.Unmarshal error: ", err)
		}
	})
	// 获取响应
	collector.OnResponse(func(response *colly.Response) {
		// 将返回中的html中所有的%去掉
		newBody := strings.ReplaceAll(string(response.Body), "%", "")
		response.Body = []byte(newBody)
	})
	// 爬取收益排行榜,(默认是按照近一年的排行)
	err := collector.Visit("https://fundact.eastmoney.com/banner/hh.html")
	if err != nil {
		global.GvaLogger.Sugar().Errorf("基金排行榜爬取失败: %s", err)
	}
}

// ConvertEntity 格式化类型
func (f *TopCrawlService) ConvertEntity() []entity.FundDayTop {
	var topList []entity.FundDayTop
	for _, item := range f.Item {
		if item.FundCode == "" {
			continue
		}
		fundTmp := entity.FundDayTop{}
		fundTmp.FundCode = item.FundCode
		// 格式化日期
		format := time.Now().Format("2006")
		fundTmp.TopDate = fmt.Sprintf("%s-%s", format, item.TopDate)
		// 转换编码
		fundTmp.FundName, _ = utils.GbkToUtf8(item.FundName)
		// 字符串转浮点型
		fundTmp.NetWorth, _ = strconv.ParseFloat(item.NetWorth, 64)
		fundTmp.DayChange, _ = strconv.ParseFloat(item.DayChange, 64)
		fundTmp.WeekChange, _ = strconv.ParseFloat(item.WeekChange, 64)
		fundTmp.MouthChange, _ = strconv.ParseFloat(item.MouthChange, 64)
		fundTmp.ThreeMouthChange, _ = strconv.ParseFloat(item.ThreeMouthChange, 64)
		fundTmp.SixMouthChange, _ = strconv.ParseFloat(item.SixMouthChange, 64)
		fundTmp.YearChange, _ = strconv.ParseFloat(item.YearChange, 64)
		fundTmp.TwoYearChange, _ = strconv.ParseFloat(item.TwoYearChange, 64)
		fundTmp.ThreeYearChange, _ = strconv.ParseFloat(item.ThreeYearChange, 64)
		fundTmp.CurrentChange, _ = strconv.ParseFloat(item.CurrentChange, 64)
		fundTmp.CreateChange, _ = strconv.ParseFloat(item.CreateChange, 64)
		topList = append(topList, fundTmp)
	}
	return topList
}
