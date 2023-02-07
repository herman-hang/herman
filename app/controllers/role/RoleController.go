package role

import (
	"github.com/gin-gonic/gin"
	BaseController "github.com/herman/app/controllers"
	RoleService "github.com/herman/app/services/role"
	RoleValidate "github.com/herman/app/validates/role"
)

func AddRole(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := BaseController.GetParams(ctx)
	RoleService.Add(RoleValidate.Add(data))
	response.Json(nil)
}
