package user

import (
	UserController "fp-back-user/app/controllers/user"
	"github.com/gin-gonic/gin"
)

// Router 用户相关路由
func Router(router *gin.RouterGroup) {
	router.POST("/login", UserController.Login)
	router.GET("/login", UserController.Login)
}
