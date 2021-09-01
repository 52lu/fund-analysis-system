package fund

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"gorm.io/gorm"
)

// 根据排行日期查找最后一条记录
func FindLastOneByDate(topDate string) (entity.FundDayTop,error)  {
	var f entity.FundDayTop
	last := global.GvaMysqlClient.Where("top_date = ?", topDate).Last(&f)
	if last.Error != nil && last.Error == gorm.ErrRecordNotFound {
		return entity.FundDayTop{}, nil
	}
	return f, last.Error
}