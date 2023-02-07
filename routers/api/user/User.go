package user

import (
	"github.com/gin-gonic/gin"
	UserController "github.com/herman/app/controllers/user"
	"github.com/herman/app/middlewares"
)

// Router 用户相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	router.Use(middlewares.Jwt("user"))
	{
		// 用户登录
		router.POST("/login", UserController.Login)
	}
}
