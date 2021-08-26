package validate

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"strings"
	"sync"
)

var (
	validate *validator.Validate        // 定义验证器变量
	chinese  = zh.New()                 // 获取中文翻译器
	uni      = ut.New(chinese, chinese) // 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")  // 获取翻译字典
)

func init()  {
	// 初始化验证器
	setValidateInstance()
}
// 获取验证器实例
func setValidateInstance() {
	var one sync.Once
	if validate == nil {
		one.Do(func() {
			// 实例化验证器
			validate = validator.New()
			// 注册自定义规则
			addRule(validate)
			// 注册翻译器
			err := zhs.RegisterDefaultTranslations(validate, trans)
			if err != nil {
				fmt.Println("RegisterDefaultTranslations-error:", err)
			}
			// 添加额外的翻译tag
			addMoreTranslationEntrance(validate, trans)
		})
	}
}

// 验证
func Validate(param interface{}) error {
	err := validate.Struct(param)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// 格式化错误
		return formatError(validationErrors)
	}
	return nil
}

/**
 * @description: 格式化错误
 * @param errs
 * @return error
 */
func formatError(errs validator.ValidationErrors) error {
	translateError := errs.Translate(trans)
	var errMsgBuffer bytes.Buffer
	for _, v := range translateError {
		//errMsgBuffer.WriteString(k)
		//errMsgBuffer.WriteString(":")
		errMsgBuffer.WriteString(strings.ToLower(v))
	}
	return errors.New(errMsgBuffer.String())
}
