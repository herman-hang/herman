package user

import (
	UserController "github.com/fp/fp-gin-framework/app/controllers/user"
	"github.com/gin-gonic/gin"
)

// Router 用户相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	// 用户登录
	router.POST("/login", UserController.Login)
}
