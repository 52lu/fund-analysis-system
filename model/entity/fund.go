package entity

import "52lu/fund-analye-system/global"

// FundBasis 基金基础信息表
type FundBasis struct {
	global.BaseModel
	Code            string  `json:"code" gorm:"type:char(6);not null; default:'';unique:un_code;comment:基金代码"`
	FullName        string  `json:"fullName" gorm:"type:varchar(50);comment:基金全称"`
	ShortName       string  `json:"shortName" gorm:"type:varchar(20);comment:基金简称"`
	MainType        string  `json:"mainType" gorm:"type:varchar(20);comment:基金主类型"`
	SubType         string  `json:"subType" gorm:"type:varchar(20);comment:基金子类型"`
	Company         string  `json:"company" gorm:"type:varchar(50);comment:基金公司"`
	Manager         string  `json:"manager" gorm:"type:varchar(20);not null; default:'';comment:基金经理人"`
	ManagerDesc     string  `json:"managerDesc" gorm:"type:varchar(255);comment:基金经理人介绍"`
	ReleaseDate     string  `json:"releaseDate" gorm:"type:varchar(12);comment:发布时间"`
	EstablishDate   string  `json:"establishDate" gorm:"type:varchar(12);comment:成立时间"`
	EstablishShares float64 `json:"establishShares" gorm:"type:decimal(12,4);comment:成立时规模(单位:亿份)"`
	ManageFeeRate   float64 `json:"manageFeeRate" gorm:"type:decimal(4,2);comment:管理费率(百分比)"`
	CustodyFeeRate  float64 `json:"custodyFeeRate" gorm:"type:decimal(4,2);comment:托管费率(百分比)"`
	SaleFeeRate     float64 `json:"saleFeeRate" gorm:"type:decimal(4,2);comment:销售服务费率(百分比)"`
	Benchmark       string  `json:"benchmark" gorm:"type:varchar(255);comment:业绩比较基准"`
}

// FundStock 基金股票持仓
type FundStock struct {
	global.BaseModel
	FundCode   string  `json:"fundCode" gorm:"type:varchar(10);not null; default:'';index;comment:基金code"`
	StockCode  string  `json:"stockCode" gorm:"type:varchar(10);index;comment:股票code"`
	Percentage float64 `json:"percentage" gorm:"type:decimal(4,2);comment:持仓占比(百分比)"`
	Quantity   float64 `json:"quantity" gorm:"type:decimal(5,2);comment:持股数(万股)"`
	Amount     float64 `json:"amount" gorm:"type:decimal(5,2);comment:持股市值(万元)"`
	CutOffDate string  `json:"cutOffDate" gorm:"type:char(10);comment:截止时间"`
}

// FundDayTop 基金每日排行
type FundDayTop struct {
	global.BaseModel
	FundCode         string  `json:"fundCode" gorm:"type:varchar(10);not null; default:'';index;comment:基金code"`
	FundName         string  `json:"fundName" gorm:"type:varchar(50);index;comment:基金名称"`
	TopDate          string  `json:"topDate" gorm:"type:varchar(12);index;comment:排名日期"`
	NetWorth         float64 `json:"netWorth" gorm:"type:decimal(10,4);comment:单位净值"`
	DayChange        float64 `json:"dayChange" gorm:"type:decimal(10,4);comment:日增长率(百分比)"`
	WeekChange       float64 `json:"weekChange" gorm:"type:decimal(10,4);comment:近一周(百分比)"`
	MouthChange      float64 `json:"mouthChange" gorm:"type:decimal(10,4);comment:近一月(百分比)"`
	ThreeMouthChange float64 `json:"threeMouthChange" gorm:"type:decimal(10,4);comment:近3个月(百分比)"`
	SixMouthChange   float64 `json:"sixMouthChange" gorm:"type:decimal(10,4);comment:近6个月(百分比)"`
	YearChange       float64 `json:"yearChange" gorm:"type:decimal(10,4);comment:近1年(百分比)"`
	TwoYearChange    float64 `json:"twoYearChange" gorm:"type:decimal(10,4);comment:近2年(百分比)"`
	ThreeYearChange  float64 `json:"threeYearChange" gorm:"type:decimal(10,4);comment:近2年(百分比)"`
	CurrentChange    float64 `json:"CurrentChange" gorm:"type:decimal(10,4);comment:今年来(百分比)"`
	CreateChange     float64 `json:"createChange" gorm:"type:decimal(10,4);comment:成立以来(百分比)"`
}
