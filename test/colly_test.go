package test

import (
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/service/crawl/fund"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
	"testing"
)

func TestFundBasis(t *testing.T) {
	fund := &fund.BaseCrawl{}
	fund.CrawlHtml("005609")
	fundEntity := &entity.FundBasis{}
	fund.ConvertToEntity(fundEntity)
	//marshal, _ := json.Marshal(fund)
	fmt.Printf("fund:%+v\n",fundEntity)
}

func TestFundTop(t *testing.T) {
	f := &fund.TopCrawl{}
	f.CrawlHtml()
	convertEntity := f.ConvertEntity()
	fmt.Printf("结果: %+v\n",convertEntity)
}



func TestReq(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnRequest(func(request *colly.Request) {

	})
	collector.OnError(func(response *colly.Response, err error) {
		t.Errorf("结果:%s\n",err)
	})
	collector.OnResponse(func(response *colly.Response) {
		var result [][]string
		body := string(response.Body)
		body = strings.Split(body,"[")[1]
		body = strings.Split(body,"]")[0]
		list := strings.Split(body,`","`)
		for _, s := range list {
			item := strings.Split(s,"|")
			result = append(result,item)
		}
		fmt.Println("结果: ",result)
	})
	err := collector.Visit("https://fundapi.eastmoney.com/fundtradenew.aspx?ft=pg&sc=r&st=desc&pi=1&pn=100&cp=&ct=&cd=&ms=&fr=&plevel=&fst=&ftype=&fr1=&fl=0&isab=1")
	if err != nil {
		t.Error(err)
	}
}