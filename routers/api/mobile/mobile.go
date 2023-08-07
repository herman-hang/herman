package mobile

import (
	"github.com/gin-gonic/gin"
	UserController "github.com/herman-hang/herman/application/controllers/admin/user"
)

// Router 移动端相关路由
// @param *gin.RouterGroup router 路由组对象
// @return void
func Router(router *gin.RouterGroup) {
	// 用户登录
	router.POST("/login", UserController.Login)
}
