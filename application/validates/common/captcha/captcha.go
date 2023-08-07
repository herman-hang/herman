package captcha

import (
	"github.com/herman-hang/herman/application/constants"
	"github.com/herman-hang/herman/application/validates"
	"github.com/herman-hang/herman/kernel/utils"
	"github.com/mitchellh/mapstructure"
)

// GetCaptchaValidate 获取验证码验证结构体
type GetCaptchaValidate struct {
	CaptchaType string `json:"captchaType" validate:"required" label:"验证码类型"`
}

// CheckCaptchaValidate 检查验证码正确性结构体
type CheckCaptchaValidate struct {
	CaptchaType string `json:"captchaType" validate:"required" label:"验证码类型"`
	Token       string `json:"token" validate:"required" label:"验证码Token"`
	PointJson   string `json:"pointJson" validate:"required" label:"验证码PointJson"`
}

// GetCaptcha 获取验证码验证
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func GetCaptcha(data map[string]interface{}) (toMap map[string]interface{}) {
	var captcha GetCaptchaValidate

	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &captcha); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(captcha); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&captcha, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}

// CheckCaptcha 获取验证码验证
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func CheckCaptcha(data map[string]interface{}) (toMap map[string]interface{}) {
	var captcha CheckCaptchaValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &captcha); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(captcha); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&captcha, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}
