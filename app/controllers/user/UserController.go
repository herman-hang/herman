package user

import (
	"github.com/gin-gonic/gin"
	BaseController "github.com/herman/app/controllers"
	UserService "github.com/herman/app/services/user"
	UserValidate "github.com/herman/app/validates/user"
)

// Login 用户登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := BaseController.GetParams(ctx)
	response.Json(UserService.Login(UserValidate.Login(data)))
}
