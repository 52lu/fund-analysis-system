/**
 * @Description 请求参数
 **/
package request

// 登录参数
type LoginParam struct {
	Phone    string `json:"phone" validate:"required,mobile"`
	Password string `json:"password" validate:"required"`
}

// 注册参数
type RegisterParam struct {
	NickName string `json:"nickName" validate:"min=2,max=4"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone" validate:"required,min=11,max=11"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
}

