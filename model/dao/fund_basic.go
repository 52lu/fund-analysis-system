package dao

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
)

type FundBasicDao struct {
}

func (d FundBasicDao) Add(entity *entity.FundBasis)  {
	global.GvaMysqlClient.Create(entity)
}
