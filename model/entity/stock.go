package entity

import "52lu/fund-analye-system/global"

// 股票基础信息
type Stock struct {
	global.BaseModel
	Code           string  `json:"code" gorm:"type:varchar(10);unique:un_code;comment:股票代码"`
	Name           string  `json:"name" gorm:"type:varchar(50);comment:股票名称"`
	Industry       string  `json:"industry" gorm:"type:varchar(20);comment:所属行业"`
	ExchangeCode   string  `json:"exchangeCode" gorm:"type:varchar(5);comment:所属交易所，SZ:深圳 SH:上海 HK:港股"`
	Tag            string  `json:"tag" gorm:"type:text;comment:所属概念"`
	SetUpDate      string  `json:"setUpDate" gorm:"type:char(10);comment:公司成立时间"`
	MarketDate     string  `json:"marketDate" gorm:"type:char(10);comment:上市时间"`
	MarketPE       float64 `json:"marketPE" gorm:"type:decimal(10,2);comment:发行市盈率"`
	MarketQuantity float64 `json:"marketQuantity" gorm:"type:decimal(10,2);comment:发行量(万股)"`
	MarketPrice    float64 `json:"marketPrice" gorm:"type:decimal(10,2);comment:发行价格"`
	MarketAmount   float64 `json:"marketAmount" gorm:"type:decimal(10,2);comment:实际募资(亿)"`
	Company        string  `json:"company" gorm:"type:varchar(255);comment:公司名称"`
	Employees      uint    `json:"employees" gorm:"comment:公司员工"`
	MoneyUnit      string  `json:"moneyUnit" gorm:"type:varchar(10);comment:货币单位"`
}

// 股票行情
type StockQuotes struct {
	global.BaseModel
	StockCode         uint    `json:"stockCode" gorm:"index;comment:股票ID"`
	RecordDate        string  `json:"recordDate" gorm:"type:char(10);comment:记录时间"`
	TotalQuantity     float64 `json:"totalQuantity" gorm:"type:decimal(10,2);comment:总股本(万股)"`
	TotalAmount       float64 `json:"totalAmount" gorm:"type:decimal(10,2);comment:总市值(亿元)"`
	CirculateQuantity float64 `json:"circulateQuantity" gorm:"type:decimal(10,2);comment:流通股本(万股)"`
	CirculateAmount   float64 `json:"circulateAmount" gorm:"type:decimal(10,2);comment:流通市值(亿)"`
	TodayOpenPrice    float64 `json:"todayOpenPrice" gorm:"type:decimal(10,2);comment:今日开盘价格"`
	TodayMaxPrice     float64 `json:"todayMaxPrice" gorm:"type:decimal(10,2);comment:今日最高价格"`
	TodayPrice        float64 `json:"todayPrice" gorm:"type:decimal(10,2);comment:今日收盘价格"`
	LastPrice         float64 `json:"lastPrice" gorm:"type:decimal(10,2);comment:昨日收盘价格"`
	MoneyUnit         string  `json:"moneyUnit" gorm:"type:varchar(10);comment:货币单位"`
	ChangeRatio       float64 `json:"changeRatio" gorm:"type:decimal(5,2);comment:换手率(百分比)"`
	DynamicPE         float64 `json:"dynamicPE" gorm:"type:decimal(10,2);comment:市盈率(动)"`
	StaticPE          float64 `json:"staticPE" gorm:"type:decimal(10,2);comment:市盈率(静)"`
	PeTtm             float64 `json:"peTTM" gorm:"type:decimal(10,2);comment:市盈率(TTM)"`
	PB                float64 `json:"PB" gorm:"type:decimal(10,2);comment:市净率"`
	EarningsPerShare  float64 `json:"earningsPerShare" gorm:"type:decimal(10,2);comment:每股收益"`
	NetAssetsPerShare float64 `json:"netAssetsPerShare" gorm:"type:decimal(10,2);comment:每股净资产"`
	Dividends         float64 `json:"dividends" gorm:"type:decimal(10,2);comment:股息(TTM)"`
	DividendsRatio    float64 `json:"dividendsRatio" gorm:"type:decimal(5,2);comment:股息率(TTM)"`
}
