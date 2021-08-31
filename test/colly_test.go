package test

import (
	"52lu/fund-analye-system/service/crawl/script"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFundBasis(t *testing.T) {
	fund := &script.FundBaseCrawl{}
	fund.GetFundBasis("005609")
	marshal, _ := json.Marshal(fund)
	fmt.Printf("fund:%s\n",marshal)
}
