/**
 * @Description TODO
 **/
package user

import (
	"52lu/fund-analye-system/global"
	"52lu/fund-analye-system/model/entity"
	"52lu/fund-analye-system/model/request"
	"gorm.io/gorm"
)

/**
 * @description: 账户密码登录
 * @param user
 */
func LoginPwd(user *entity.User) error {
	//校验账户和密码
	result := global.GvaMysqlClient.Where("phone=? and password=?", user.Phone, user.Password).
		First(user)
	if result.Error != nil {
		return result.Error
	}
	// 查询用户信息
	userInfo := entity.UserInfo{}
	result = global.GvaMysqlClient.Where("uid = ?", user.ID).First(&userInfo)
	if result.Error != nil {
		return result.Error
	}
	user.UserInfo = userInfo
	return result.Error
}

// 注册用户
func Register(param request.RegisterParam) (*entity.User, error) {
	user := entity.User{
		NickName: param.NickName,
		Phone:    param.Phone,
		Password: param.Password,
	}
	err := global.GvaMysqlClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			global.GvaLogger.Sugar().Errorf("新增用户失败: %s", err)
			return err
		}
		userInfo := entity.UserInfo{
			Uid:      user.ID,
			Birthday: param.Birthday,
			Address:  param.Address,
		}
		if err := tx.Create(&userInfo).Error; err != nil {
			global.GvaLogger.Sugar().Errorf("新增用户信息失败: %s", err)
			return err
		}
		user.UserInfo = userInfo
		return nil
	})
	//user.Token =
	return &user, err
}
