package captcha

import (
	"fmt"
	CaptchaConstant "github.com/herman-hang/herman/application/constants/common/captcha"
	"github.com/herman-hang/herman/kernel/utils"
)

// GetCaptcha 获取验证码
// @param map data 前端请求数据
// @return captchaData 返回验证码相关信息
func GetCaptcha(data map[string]interface{}) (captchaData map[string]interface{}) {
	captchaData, err := utils.Factory().GetService(data["captchaType"].(string)).Get()
	if err != nil {
		panic(CaptchaConstant.GetCaptchaFail)
	}
	return captchaData
}

// CheckCaptcha 检查验证码
// @param map data 前端请求数据
// @return captchaData 返回验证码相关信息
func CheckCaptcha(data map[string]interface{}) {
	err := utils.Factory().GetService(fmt.Sprintf("%s", data["captchaType"])).
		Check(fmt.Sprintf("%s", data["token"]),
			fmt.Sprintf("%s", data["pointJson"]))

	if err != nil {
		panic(CaptchaConstant.CheckCaptchaError)
	}
}
