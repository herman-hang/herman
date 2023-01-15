package admin

import (
	BaseController "github.com/fp/fp-gin-framework/app/controllers"
	AdminService "github.com/fp/fp-gin-framework/app/services/admin"
	AdminValidate "github.com/fp/fp-gin-framework/app/validates/admin"
	"github.com/gin-gonic/gin"
)

// Login 管理员登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := BaseController.GetParams(ctx)
	response.Json(AdminService.Login(AdminValidate.Login(data)))
}
