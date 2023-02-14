package user

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	UserService "github.com/herman/app/services/user"
	UserValidate "github.com/herman/app/validates/user"
)

// Login 用户登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(UserService.Login(UserValidate.Login(data)))
}
