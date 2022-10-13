package user

import (
	UserController "fp-back-user/app/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router 用户相关路由
func Router(router *gin.RouterGroup) {
	router.POST("/login", UserController.Login)
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "测试! ",
		})
	})
}
