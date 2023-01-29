package admin

import (
	AdminController "github.com/fp/fp-gin-framework/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

// Router 管理员相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	// 管理员登录
	router.POST("/login", AdminController.Login)
}
