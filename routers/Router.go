package routers

import (
	"github.com/fp/fp-gin-framework/app/middlewares"
	"github.com/fp/fp-gin-framework/routers/api/user"
	"github.com/fp/fp-gin-framework/servers/settings"
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
