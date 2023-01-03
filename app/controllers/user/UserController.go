package user

import (
	"github.com/fp/fp-gin-framework/app"
	base "github.com/fp/fp-gin-framework/app/controllers"
	userService "github.com/fp/fp-gin-framework/app/services/user"
	userValidate "github.com/fp/fp-gin-framework/app/validates/user"
	"github.com/gin-gonic/gin"
)

// Login 用户列表
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := base.GetParams(ctx)
	// Success参数可以设置零个或多个
	response.Success(app.D(userService.Login(userValidate.Login(data))))
}
