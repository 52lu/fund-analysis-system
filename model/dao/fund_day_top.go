package dao

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"gorm.io/gorm"
)

// 根据排行日期查找最后一条记录
func FindLastOneByDate(topDate string) (entity.FundDayTop, error) {
	var f entity.FundDayTop
	last := global.GvaMysqlClient.Where("top_date = ?", topDate).Last(&f)
	if last.Error != nil && last.Error == gorm.ErrRecordNotFound {
		return entity.FundDayTop{}, nil
	}
	return f, last.Error
}

type FilterBasicResult struct {
	FundCode string
	Code     string
}

// 查询没有详情的基金信息
func FilterBasicFund() []FilterBasicResult {
	res := []FilterBasicResult{}
	global.GvaMysqlClient.Raw("SELECT A.fund_code,B.`code` from fas_fund_day_top as A  LEFT JOIN fas_fund_basis as B on  A.fund_code = B.`code` WHERE  B.`code` is NULL GROUP BY A.fund_code").Scan(&res)
	return res
}

