package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	CaptchaController "github.com/herman/app/controllers/captcha"
	"github.com/herman/routers/api/admin"
	"github.com/herman/routers/api/user"
	"github.com/herman/servers/settings"
)

// InitRouter 初始化路由
// @param *gin.Engine rootEngine 路由引擎
// @return void
func InitRouter(rootEngine *gin.Engine) {
	// 测试路由
	rootEngine.GET("/", func(context *gin.Context) {
		response := app.Request{Context: context}
		response.Success(app.D(map[string]interface{}{
			"welcome": "Hello fp-gin-framework!",
		}))
	})
	// 设置路由前缀
	api := rootEngine.Group(settings.Config.AppPrefix)
	// 获取验证码
	api.GET("/captcha", CaptchaController.GetCaptcha)
	// 检查验证码正确性
	api.POST("/captcha/check", CaptchaController.CheckCaptcha)

	// 用户模块
	userRouter := api.Group("/user")
	// 后台模块
	adminRouter := api.Group("/admin")

	user.Router(userRouter)
	admin.Router(adminRouter)
}
