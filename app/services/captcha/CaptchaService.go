package captcha

import (
	"fmt"
	"github.com/fp/fp-gin-framework/app/common"
)

// Captcha 获取验证码
// @param map data 前端请求数据
// @return captchaData 返回验证码相关信息
func Captcha(data map[string]interface{}) (captchaData map[string]interface{}) {
	captchaData, err := common.Captcha.GetService(fmt.Sprintf("%s", data["captcha_type"])).Get()
	if err != nil {
		return nil
	}
	return captchaData
}
