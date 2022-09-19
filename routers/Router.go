package routers

import (
	"fp-back-user/app/middlewares"
	"fp-back-user/routers/api/user"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(rootEngine *gin.Engine) {
	api := rootEngine.Group("/api/v1")

	api.Use(middlewares.Jwt())
	{
		// 用户相关路由
		userRouter := api.Group("/user")

		user.Router(userRouter)
	}
}
