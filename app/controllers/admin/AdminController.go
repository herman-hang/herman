package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/herman/app"
	AdminService "github.com/herman/app/services/admin"
	AdminValidate "github.com/herman/app/validates/admin"
)

// Login 管理员登录
// @param *gin.Context ctx 上下文
// @return json
func Login(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Login(AdminValidate.Login(data)))
}

// AddAdmin 管理员添加
// @param *gin.Context ctx 上下文
// @return json
func AddAdmin(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	AdminService.Add(AdminValidate.Add.Check(data))
	context.Json(nil)
}

// ModifyAdmin 管理员修改
// @param *gin.Context ctx 上下文
// @return json
func ModifyAdmin(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	AdminService.Modify(AdminValidate.Modify.Check(data))
	context.Json(nil)
}

// FindAdmin 根据ID查询管理员详情
// @param *gin.Context ctx 上下文
// @return json
func FindAdmin(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminService.Find(AdminValidate.Find.Check(data)))
}
