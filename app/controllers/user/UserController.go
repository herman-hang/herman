package user

import (
	"fp-back-user/app"
	base "fp-back-user/app/controllers"
	userService "fp-back-user/app/services/user"
	userValidate "fp-back-user/app/validates/user"
	"github.com/gin-gonic/gin"
)

// Login 用户列表
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	this, data := base.GetParams(ctx)
	// Response参数可以设置零个或多个
	this.Response(app.D(userService.Login(userValidate.Login(data))))
	return
}
