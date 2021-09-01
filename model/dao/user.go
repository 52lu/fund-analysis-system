package dao

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
)

type UserDao struct {
	Uid uint
}

// FindUser 查询用户信息
func (u *UserDao) FindUser() (*entity.User, error) {
	var user entity.User
	//校验账户和密码
	result := global.GvaMysqlClient.Where("id=? ", u.Uid).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 查询用户信息
	userInfo := entity.UserInfo{}
	result = global.GvaMysqlClient.Where("uid = ?", u.Uid).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	user.UserInfo = userInfo
	return &user, nil
}
