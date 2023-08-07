package user

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	UserService "github.com/herman-hang/herman/application/services/admin/user"
	UserValidate "github.com/herman-hang/herman/application/validates/admin/user"
)

// Login 用户登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(UserService.Login(UserValidate.Login(data)))
}
