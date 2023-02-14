package role

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	RoleService "github.com/herman/app/services/role"
	RoleValidate "github.com/herman/app/validates/role"
)

// AddRole 添加角色
// @param *gin.Context ctx 上下文
// @return void
func AddRole(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	RoleService.Add(RoleValidate.Add(data))
	context.Json(nil)
}

func ModifyRole(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	context.Json(nil)
}

func FindRole(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	context.Json(nil)
}

func DeleteRole(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	context.Json(nil)
}

func ListRole(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	context.Json(nil)
}
