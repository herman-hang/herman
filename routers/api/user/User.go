package user

import (
	"fp-back-user/app/controllers"
	"github.com/gin-gonic/gin"
)

// Router 用户相关路由
func Router(r *gin.RouterGroup) {
	r.POST("/login", controllers.UserLogin)
}
