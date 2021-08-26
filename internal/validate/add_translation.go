// validate: 翻译额外字段
package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// 添加翻译器
func _registerTranslator(tag,msg string) validator.RegisterTranslationsFunc  {
	return func(ut ut.Translator) error {
		if err := ut.Add(tag,msg,false); err != nil {
			return err
		}
		return nil
	}
}
func _translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}
	return t
}

// 添加额外翻译Tag入口
func addMoreTranslationEntrance(validate *validator.Validate,translator ut.Translator)  {
	// 添加手机号错误翻译
	addMobileTrans(validate,translator)
}
// 添加手机号错误翻译
func addMobileTrans(validate *validator.Validate,translator ut.Translator)  {
	_ = validate.RegisterTranslation("mobile", translator,
		_registerTranslator("mobile", "{0}格式不正确！"),
		_translateFunc)
}