package user

import (
	"fp-back-user/app"
	base "fp-back-user/app/controllers"
	userService "fp-back-user/app/services/user"
	userValidate "fp-back-user/app/validates/user"
	"github.com/gin-gonic/gin"
)

// Login 用户列表
// controller中只负责接收数据，数据验证，函数调用，返回数据，不做其他业务处理，业务实现全部要在service中
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	this, data := base.GetParams(ctx)

	// Response参数可以设置零个或多个
	this.Response(
		app.D(userService.Login(userValidate.Login(data))),
	)
	return
}
