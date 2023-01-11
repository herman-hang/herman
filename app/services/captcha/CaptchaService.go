package captcha

import (
	"fmt"
	captchaConstants "github.com/fp/fp-gin-framework/app/constants/captcha"
	"github.com/fp/fp-gin-framework/app/utils"
)

// GetCaptcha 获取验证码
// @param map data 前端请求数据
// @return captchaData 返回验证码相关信息
func GetCaptcha(data map[string]interface{}) (captchaData map[string]interface{}) {
	captchaData, err := utils.Factory().GetService(fmt.Sprintf("%s", data["captchaType"])).Get()
	if err != nil {
		panic(captchaConstants.GetCaptchaFail)
	}
	return captchaData
}

// CheckCaptcha 获取验证码
// @param map data 前端请求数据
// @return captchaData 返回验证码相关信息
func CheckCaptcha(data map[string]interface{}) (err error) {
	err = utils.Factory().GetService(fmt.Sprintf("%s", data["captchaType"])).
		Check(fmt.Sprintf("%s", data["Token"]),
			fmt.Sprintf("%s", data["PointJson"]))
	if err != nil {
		panic(captchaConstants.CheckCaptchaError)
	}

	return nil
}
