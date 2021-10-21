package fund

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/dao"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/service/crawl"
	"52lu/fund-analye-system/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

//   定义结构体对应
type BasisCrawl struct {
	Code            string `selector:"tr:nth-child(2) > td:nth-of-type(1)"`
	FullName        string `selector:"tr:nth-child(1) > td:nth-of-type(1)"`
	ShortName       string `selector:"tr:nth-child(1) > td:nth-of-type(2)"`
	Type            string `selector:"tr:nth-child(2) > td:nth-of-type(2)"`
	ReleaseDate     string `selector:"tr:nth-child(3) > td:nth-of-type(1)"`
	EstablishDate   string `selector:"tr:nth-child(3) > td:nth-of-type(2)"`
	EstablishShares string `selector:"tr:nth-child(3) > td:nth-of-type(2)"`
	Company         string `selector:"tr:nth-child(5) > td:nth-of-type(1)"`
	Manager         string `selector:"tr:nth-child(6) > td:nth-of-type(1)"`
	ManagerDesc     string `selector:"tr:nth-child(6) > td:nth-of-type(1) > a[href]" attr:"href"`
	ManageFeeRate   string `selector:"tr:nth-child(7) > td:nth-of-type(1)"`
	CustodyFeeRate  string `selector:"tr:nth-child(7) > td:nth-of-type(2)"`
	SaleFeeRate     string `selector:"tr:nth-child(8) > td:nth-of-type(1)"`
	Benchmark       string `selector:"tr:nth-child(10) > td:nth-of-type(1)"`
}

// 分组爬取
func splitFundBasicList(data []dao.FilterBasicResult, groupNum int) [][]dao.FilterBasicResult {
	var result [][]dao.FilterBasicResult
	length := len(data)
	if len(data) <= groupNum {
		return append(result, data)
	}
	perGroupNum := int(math.Ceil(float64(length) / float64(groupNum)))
	index := 0
	// 分groupNum组
	for i := 0; i < groupNum; i++ {
		// 每组数量是perGroupNum
		group := []dao.FilterBasicResult{}
		for j := 0; j < perGroupNum; j++ {
			if index <= length-1 {
				group = append(group, data[index])
				index++
			}
		}
		result = append(result, group)
	}
	return result
}

// 批量抓取
func BatchBasicCrawl() {
	// 从排行榜中获取code,并过滤已经爬取过的code
	basicFundList := dao.FilterBasicFund()
	total := len(basicFundList)
	if total > 0 {
		var baseRowsChannel = make(chan entity.FundBasis, total)
		// 分组抓取
		crawlByGroup(basicFundList, baseRowsChannel)
		// 遍历channel获取数据
		var fundBasisRows []entity.FundBasis
		for item := range baseRowsChannel {
			fundBasisRows = append(fundBasisRows, item)
		}
		if fundBasisRows != nil {
			// 保存入库
			create := global.GvaMysqlClient.Create(fundBasisRows)
			if create.Error != nil {
				global.GvaLogger.Sugar().Errorf("基金详情入库失败", create.Error)
				return
			}
			global.GvaLogger.Sugar().Infof("基金详情抓取成功，共: %v 条", create.RowsAffected)
		}
	}
}

// 分组抓取，防止并发过大，被拒绝访问
func crawlByGroup(basicResults []dao.FilterBasicResult, c chan<- entity.FundBasis) {
	// 分组抓取
	groupNum := 15
	fundCodeGroup := splitFundBasicList(basicResults, groupNum)
	// 并发请求抓取
	var wg sync.WaitGroup
	wg.Add(groupNum)
	for _, results := range fundCodeGroup {
		basicFundList := results
		go func() {
			for _, item := range basicFundList {
				filterBasicResult := item
				f := BasisCrawl{}
				// 爬取页面信息
				f.CrawlHtml(filterBasicResult.FundCode)
				if f.Code != "" {
					// 转成实体类型
					toEntity := f.ConvertToEntity()
					c <- toEntity
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// 关闭通道
	close(c)
}

//  抓取单个基金基本信息
func (f *BasisCrawl) CrawlHtml(fundCode string) {
	collector := colly.NewCollector(colly.UserAgent(crawl.UserAgent), colly.Async(true))
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
	// 开启限速
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*fundf10.eastmoney.*",
		Delay:       500 * time.Millisecond,
		RandomDelay: 500 * time.Millisecond,
		Parallelism: 20,
	})
	if err != nil {
		global.GvaLogger.Sugar().Errorf("设置限速失败: %s", err)
		return
	}
	err = collector.Visit(fmt.Sprintf("https://fundf10.eastmoney.com/jbgk_%s.html", fundCode))
	if err != nil {
		global.GvaLogger.Sugar().Errorf("基金%s,信息请求失败: %s", fundCode, err)
	}
	collector.Wait()
}

// ConvertToEntity 格式化数据为实体类
func (f *BasisCrawl) ConvertToEntity() entity.FundBasis {
	if f.Code == "" {
		return entity.FundBasis{}
	}
	var fundBaseEntity entity.FundBasis
	// 部分基金code解析为: 006049（前端）、006050（后端）,如：https://fundf10.eastmoney.com/jbgk_006049.html
	if strings.Contains(f.Code, "、") {
		f.Code = strings.Split(f.Code, "、")[0]
	}
	fundBaseEntity.Code = utils.ExtractNumberFromString(f.Code)
	fundBaseEntity.FullName = f.FullName
	fundBaseEntity.ShortName = f.ShortName
	// 类型分割
	typeInfo := strings.Split(f.Type, "-")
	fundBaseEntity.MainType = typeInfo[0]
	fundBaseEntity.SubType = typeInfo[1]
	// 基金公司
	fundBaseEntity.Company = f.Company
	// 基金经理
	fundBaseEntity.Manager = f.Manager
	fundBaseEntity.ManagerDesc = strings.ReplaceAll(f.ManagerDesc, "//", "")
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
	return fundBaseEntity
}

// 处理日期，把年、月、日替换成-
func replaceDateChinese(strDate string) string {
	strDate = strings.ReplaceAll(strDate, "年", "-")
	strDate = strings.ReplaceAll(strDate, "月", "-")
	strDate = strings.ReplaceAll(strDate, "日", "")
	return strDate
}
