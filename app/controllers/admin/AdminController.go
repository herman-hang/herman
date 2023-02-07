package admin

import (
	"github.com/gin-gonic/gin"
	BaseController "github.com/herman/app/controllers"
	AdminService "github.com/herman/app/services/admin"
	AdminValidate "github.com/herman/app/validates/admin"
)

// Login 管理员登录
// @param *gin.Context ctx 上下文
func Login(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := BaseController.GetParams(ctx)
	response.Json(AdminService.Login(AdminValidate.Login(data)))
}
