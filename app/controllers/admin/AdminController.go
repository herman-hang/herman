package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	AdminService "github.com/herman/app/services/admin"
	AdminValidate "github.com/herman/app/validates/admin"
)

// Login 管理员登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Login(AdminValidate.Login(data)))
}
