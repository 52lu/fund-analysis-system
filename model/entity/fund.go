package entity

import "52lu/fund-analye-system/global"

// 基金基础信息表
type FundBasis struct {
	global.BaseModel
	Code            string  `json:"code" gorm:"type:char(6);unique:un_code;comment:基金代码"`
	FullName        string  `json:"fullName" gorm:"type:varchar(50);comment:基金全称"`
	ShortName       string  `json:"shortName" gorm:"type:varchar(50);comment:基金简称"`
	Type            string  `json:"type" gorm:"type:varchar(50);comment:基金类型"`
	Company         string  `json:"company" gorm:"type:varchar(50);comment:基金公司"`
	ReleaseDate     string  `json:"releaseDate" gorm:"type:varchar(12);comment:发布时间"`
	EstablishDate   string  `json:"establishDate" gorm:"type:varchar(12);comment:成立时间"`
	EstablishShares float64 `json:"establishShares" gorm:"type:decimal(12,4);comment:成立时规模(单位:亿份)"`
	ManageFeeRate   float64 `json:"manageFeeRate" gorm:"type:decimal(4,2);comment:管理费率(百分比)"`
	CustodyFeeRate  float64 `json:"custodyFeeRate" gorm:"type:decimal(4,2);comment:托管费率(百分比)"`
	SaleFeeRate     float64 `json:"saleFeeRate" gorm:"type:decimal(4,2);comment:销售服务费率(百分比)"`
	Benchmark       float64 `json:"benchmark" gorm:"type:varchar(255);comment:业绩比较基准"`
}

// 基金股票持仓
type FundStock struct {
	global.BaseModel
	FundCode   string  `json:"fundCode" gorm:"type:varchar(10);index;comment:基金code"`
	StockCode  string  `json:"stockCode" gorm:"type:varchar(10);index;comment:股票code"`
	Percentage float64 `json:"percentage" gorm:"type:decimal(4,2);comment:持仓占比(百分比)"`
	Quantity   float64 `json:"quantity" gorm:"type:decimal(5,2);comment:持股数(万股)"`
	Amount     float64 `json:"amount" gorm:"type:decimal(5,2);comment:持股市值(万元)"`
	CutOffDate string  `json:"cutOffDate" gorm:"type:char(10);comment:截止时间"`
}
