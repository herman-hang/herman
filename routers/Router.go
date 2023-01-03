package routers

import (
	"github.com/fp/fp-gin-framework/app/constants"
	captchaController "github.com/fp/fp-gin-framework/app/controllers/captcha"
	"github.com/fp/fp-gin-framework/app/middlewares"
	"github.com/fp/fp-gin-framework/routers/api/user"
	"github.com/fp/fp-gin-framework/servers/settings"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
// @param *gin.Engine rootEngine 路由引擎
// @return void
func InitRouter(rootEngine *gin.Engine) {
	// 测试路由
	rootEngine.GET("/", func(context *gin.Context) {
		context.JSON(constants.SuccessCode, gin.H{
			"code":    constants.SuccessCode,
			"message": constants.Success,
			"data":    nil,
		})
	})
	// 设置路由前缀
	api := rootEngine.Group(settings.Config.AppPrefix)
	// 获取验证码
	api.GET("/captcha", captchaController.GetCaptcha)

	api.Use(middlewares.Jwt("user"))
	{
		// 用户相关路由
		userRouter := api.Group("/user")

		user.Router(userRouter)
	}
}
