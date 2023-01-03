package captcha

import (
	"github.com/fp/fp-gin-framework/app/constants"
	captchaConstant "github.com/fp/fp-gin-framework/app/constants/captcha"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/fp/fp-gin-framework/app/validates"
	"github.com/mitchellh/mapstructure"
)

type GetCaptchaValidate struct {
	CaptchaType int    `json:"captcha_type" validate:"required,numeric,oneof=1 2" label:"验证码类型"`
	Token       string `json:"token"`
	PointJson   string `json:"point_json"`
}

// Captcha 获取验证码验证
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func Captcha(data map[string]interface{}) (toMap map[string]interface{}) {
	var (
		captcha     GetCaptchaValidate
		CaptchaType = map[int]string{
			captchaConstant.BlockPuzzle: "blockPuzzle",
			captchaConstant.ClickWord:   "clickWord",
		}
	)

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

	// 从interface{}转为int类型
	toMap["captcha_type"] = CaptchaType[captcha.CaptchaType]
	return toMap
}
