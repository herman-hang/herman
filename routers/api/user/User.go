package user

import (
	UserController "fp-back-user/app/controllers/user"
	"github.com/gin-gonic/gin"
)

// Router 用户相关路由
// @param *gin.RouterGroup router // 路由组对象
func Router(router *gin.RouterGroup) {
	router.POST("/login", UserController.Login)
}
