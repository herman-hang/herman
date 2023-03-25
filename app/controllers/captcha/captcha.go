package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	CaptchaService "github.com/herman-hang/herman/app/services/captcha"
	CaptchaValidate "github.com/herman-hang/herman/app/validates/captcha"
)

// GetCaptcha 获取验证码（支持2钟验证码，请求参数CaptchaType为1：滑动拼图，CaptchaType为2：文字点选）
// @param *gin.Context ctx 上下文
// @return void
func GetCaptcha(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(CaptchaService.GetCaptcha(CaptchaValidate.GetCaptcha(data)))
}

// CheckCaptcha 检查验证码正确性
// @param *gin.Context ctx 上下文
// @return void
func CheckCaptcha(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(CaptchaService.CheckCaptcha(CaptchaValidate.CheckCaptcha(data)))
}
