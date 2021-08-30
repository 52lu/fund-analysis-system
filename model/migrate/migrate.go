/**
 * @Description mysql迁移
 **/
package migrate

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"fmt"
	"gorm.io/gorm"
)

// 设置表信息
func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GvaMysqlClient.Set("gorm:table_options", value)
}

// 数据表迁移
func AutoMigrate() {
	// 用户相关表
	userMigrate()
	// 基金相关表
	fundMigrate()
	// 股票相关表
	stockMigrate()
}

// ################ 具体表 #########################

// 用户相关表
func userMigrate() {
	// 用户账号表
	_ = setTableOption("用户表").AutoMigrate(&entity.User{})
	// 用户信息表
	_ = setTableOption("用户信息表").AutoMigrate(&entity.UserInfo{})
}
// 基金表
func fundMigrate()  {
	// 基金基础表
	_ = setTableOption("基金表").AutoMigrate(&entity.FundBasis{})
	// 基金持仓表
	_ = setTableOption("基金持仓表").AutoMigrate(&entity.FundStock{})
}

// 股票相关表
func stockMigrate()  {
	// 股票基础表
	_ = setTableOption("股票表").AutoMigrate(&entity.Stock{})
	// 股票行情表
	_ = setTableOption("股票行情表").AutoMigrate(&entity.StockQuotes{})
}