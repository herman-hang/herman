package role

import (
	"github.com/gin-gonic/gin"
	BaseController "github.com/herman/app/controllers"
	RoleService "github.com/herman/app/services/role"
	RoleValidate "github.com/herman/app/validates/role"
)

// AddRole 添加角色
// @param *gin.Context ctx 上下文
// @return void
func AddRole(ctx *gin.Context) {
	// 接收gin上下文和请求数据
	data, response := BaseController.GetParams(ctx)
	RoleService.Add(RoleValidate.Add(data))
	response.Json(nil)
}

func ModifyRole(ctx *gin.Context) {
	_, response := BaseController.GetParams(ctx)
	response.Json(nil)
}

func FindRole(ctx *gin.Context) {
	_, response := BaseController.GetParams(ctx)
	response.Json(nil)
}

func DeleteRole(ctx *gin.Context) {
	_, response := BaseController.GetParams(ctx)
	response.Json(nil)
}

func ListRole(ctx *gin.Context) {
	_, response := BaseController.GetParams(ctx)
	response.Json(nil)
}
