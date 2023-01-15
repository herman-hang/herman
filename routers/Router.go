package routers

import (
	"github.com/fp/fp-gin-framework/app"
	CaptchaController "github.com/fp/fp-gin-framework/app/controllers/captcha"
	"github.com/fp/fp-gin-framework/app/middlewares"
	"github.com/fp/fp-gin-framework/routers/api/admin"
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
		response := app.Request{Context: context}
		response.Success(app.D(map[string]interface{}{
			"test": "Hello fp-gin-framework!",
		}))
	})
	// 设置路由前缀
	api := rootEngine.Group(settings.Config.AppPrefix)
	// 获取验证码
	api.GET("/captcha", CaptchaController.GetCaptcha)
	// 检查验证码正确性
	api.POST("/captcha/check", CaptchaController.CheckCaptcha)

	// 用户模块
	api.Use(middlewares.Jwt("user"))
	{
		userRouter := api.Group("/user")

		user.Router(userRouter)
	}

	// 管理员模块
	api.Use(middlewares.Jwt("admin"))
	{
		adminRouter := api.Group("/admin")

		admin.Router(adminRouter)
	}
}
