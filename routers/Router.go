package routers

import (
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
		context.JSON(200, gin.H{
			"code":    200,
			"message": "操作成功",
			"data":    nil,
		})
	})
	// 设置路由前缀
	api := rootEngine.Group(settings.Config.AppPrefix)
	api.Use(middlewares.Jwt())
	{
		// 用户相关路由
		userRouter := api.Group("/user")

		user.Router(userRouter)
	}
}
