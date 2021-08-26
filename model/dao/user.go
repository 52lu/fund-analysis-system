/**
 * @Description 操作用户表
 **/
package dao

import (
	"52lu/fund-analye-system/global"
	userEntity "52lu/fund-analye-system/model/entity/user"
)

type UserDao struct {
	Uid uint
}

// 查询用户信息
func (u *UserDao) FindUser() (*userEntity.User, error) {
	var user userEntity.User
	//校验账户和密码
	result := global.GvaMysqlClient.Where("id=? ", u.Uid).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 查询用户信息
	userInfo := userEntity.UserInfo{}
	result = global.GvaMysqlClient.Where("uid = ?", u.Uid).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	user.UserInfo = userInfo
	return &user, nil
}
