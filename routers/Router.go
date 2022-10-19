package routers

import (
	"fp-back-user/app/middlewares"
	"fp-back-user/routers/api/user"
	"fp-back-user/settings"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
// @param *gin.Engine rootEngine 路由引擎
func InitRouter(rootEngine *gin.Engine) {
	api := rootEngine.Group(settings.Config.AppPrefix)

	api.Use(middlewares.Jwt())
	{
		// 用户相关路由
		userRouter := api.Group("/user")

		user.Router(userRouter)
	}
}
