package fund

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
)

// 根据排行日期查找最后一条记录
func FindLastOneByDate(topDate string) (entity.FundDayTop,error)  {
	var f entity.FundDayTop
	last := global.GvaMysqlClient.Where("top_date = ?", topDate).Last(&f)
	return entity.FundDayTop{}, last.Error
}