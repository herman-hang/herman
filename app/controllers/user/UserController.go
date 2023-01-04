package user

import (
	base "github.com/fp/fp-gin-framework/app/controllers"
	userService "github.com/fp/fp-gin-framework/app/services/user"
	userValidate "github.com/fp/fp-gin-framework/app/validates/user"
	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := base.GetParams(ctx)
	response.Json(userService.Login(userValidate.Login(data)))
}
