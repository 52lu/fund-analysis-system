// validate: 添加自定义tag
package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 注册自定义规则
func addRule(v *validator.Validate) {
	_ = v.RegisterValidation("mobile", addMobileTag)
}

// 注册手机号验证规则
func addMobileTag(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}
