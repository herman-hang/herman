package system

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	SystemService "github.com/herman-hang/herman/application/services/admin/system"
	SystemValidate "github.com/herman-hang/herman/application/validates/admin/system"
)

// FindSystem 获取系统设置信息
// @param *gin.Context ctx 上下文
// @return void
func FindSystem(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	context.Json(SystemService.Find())
}

// ModifySystem 修改系统设置信息
// @param *gin.Context ctx 上下文
// @return void
func ModifySystem(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	SystemService.Modify(SystemValidate.Modify.Check(data))
	context.Json(nil)
}
