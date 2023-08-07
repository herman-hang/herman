package role

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	RoleService "github.com/herman-hang/herman/application/services/admin/role"
	RoleValidate "github.com/herman-hang/herman/application/validates/admin/role"
)

// AddRole 添加角色
// @param *gin.Context ctx 上下文
// @return void
func AddRole(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	RoleService.Add(RoleValidate.Add.Check(data))
	context.Json(nil)
}

// ModifyRole 修改角色
// @param *gin.Context ctx 上下文
// @return void
func ModifyRole(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	RoleService.Modify(RoleValidate.Modify.Check(data))
	context.Json(nil)
}

// FindRole 根据ID获取角色详情
// @param *gin.Context ctx 上下文
// @return void
func FindRole(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(RoleService.Find(RoleValidate.Find.Check(data)))
}

// RemoveRole 删除角色
// @param *gin.Context ctx 上下文
// @return void
func RemoveRole(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	RoleService.Remove(RoleValidate.Delete.Check(data))
	context.Json(nil)
}

// ListRole 角色列表
// @param *gin.Context ctx 上下文
// @return void
func ListRole(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	RoleService.List(data)
	context.Json(RoleService.List(RoleValidate.List.Check(data)))
}
