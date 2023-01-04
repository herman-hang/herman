package captcha

import (
	base "github.com/fp/fp-gin-framework/app/controllers"
	captchaService "github.com/fp/fp-gin-framework/app/services/captcha"
	captchaValidate "github.com/fp/fp-gin-framework/app/validates/captcha"
	"github.com/gin-gonic/gin"
)

// GetCaptcha 获取验证码（支持2钟验证码，请求参数CaptchaType为1：滑动拼图，CaptchaType为2：文字点选）
// @param *gin.Context ctx 上下文
func GetCaptcha(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := base.GetParams(ctx)
	// 响应操作
	response.Json(captchaService.GetCaptcha(captchaValidate.GetCaptcha(data)))
}

// CheckCaptcha 检查验证码正确性
// @param *gin.Context ctx 上下文
func CheckCaptcha(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := base.GetParams(ctx)
	// 响应操作
	response.Json(captchaService.CheckCaptcha(captchaValidate.CheckCaptcha(data)))
}
