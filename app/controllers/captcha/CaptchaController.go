package captcha

import (
	"github.com/fp/fp-gin-framework/app"
	base "github.com/fp/fp-gin-framework/app/controllers"
	captchaService "github.com/fp/fp-gin-framework/app/services/captcha"
	captchaValidate "github.com/fp/fp-gin-framework/app/validates/captcha"
	"github.com/gin-gonic/gin"
)

// GetCaptcha 获取验证码
// @param *gin.Context ctx 上下文
func GetCaptcha(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := base.GetParams(ctx)
	// Success参数可以设置零个或多个
	response.Success(app.D(captchaService.Captcha(captchaValidate.Captcha(data))))
}
