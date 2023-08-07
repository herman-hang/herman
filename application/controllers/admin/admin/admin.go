package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	AdminConstant "github.com/herman-hang/herman/application/constants/admin"
	AdminService "github.com/herman-hang/herman/application/services/admin/admin"
	"github.com/herman-hang/herman/application/validates/admin/admin"
)

// Login 管理员登录
// @param *gin.Context ctx 上下文
// @return void
func Login(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Login(admin.Login(data), ctx), AdminConstant.LoginSuccess)
}

// AddAdmin 管理员添加
// @param *gin.Context ctx 上下文
// @return void
func AddAdmin(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	AdminService.Add(admin.Add.Check(data))
	context.Json(nil)
}

// ModifyAdmin 管理员修改
// @param *gin.Context ctx 上下文
// @return void
func ModifyAdmin(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	AdminService.Modify(admin.Modify.Check(data))
	context.Json(nil)
}

// FindAdmin 根据ID查询管理员详情
// @param *gin.Context ctx 上下文
// @return void
func FindAdmin(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Find(admin.Find.Check(data)))
}

// RemoveAdmin 删除管理员
// @param *gin.Context ctx 上下文
// @return void
func RemoveAdmin(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	AdminService.Remove(admin.Delete.Check(data))
	context.Json(nil)
}

// ListAdmin 管理员列表
// @param *gin.Context ctx 上下文
// @return void
func ListAdmin(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.List(admin.List.Check(data)))
}
