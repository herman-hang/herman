package validates

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// Validate 全局model数据验证器
// @param 接收一个待数据验证的结构体
// @return err 返回错误信息
func Validate(dataStruct interface{}) (err error) {
	// 验证
	zhCh := zh.New()
	validate := validator.New()

	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	// 验证器注册翻译器
	err = zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}

	err = validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}

	return nil
}
