/**
 * @Description 用户相关的实体
 **/
package user

import (
	"52lu/fund-analye-system/global"
)

// 用户表
type User struct {
	global.BaseModel
	NickName string   `json:"nickName" gorm:"type:varchar(20);not null;default:'';comment:昵称"`
	Phone    string   `json:"phone" gorm:"type:char(11);unique:un_phone;comment:手机号"`
	Password string   `json:"-" gorm:"type:varchar(40);comment:密码"`
	Status   int      `json:"status" gorm:"size:4;default:1;comment:状态 1:正常 2:白名单 3:黑名单"`
	UserInfo UserInfo `json:"userInfo" gorm:"-"`
	Token    string   `json:"token" gorm:"-"`
}

// 用户信息表
type UserInfo struct {
	global.BaseModel
	Uid      uint   `json:"uid" gorm:"comment:用户id"`
	Birthday string `json:"birthday" gorm:"type:varchar(10);comment:生日"`
	Address  string `json:"address" gorm:"type:text;comment:地址"`
}
