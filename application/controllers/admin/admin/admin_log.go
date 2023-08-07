package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	AdminLogService "github.com/herman-hang/herman/application/services/admin/admin"
	AdminValidate "github.com/herman-hang/herman/application/validates/admin/admin"
)

// LogList 管理员日志列表
// @param *gin.Context ctx 上下文
// @return void
func LogList(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminLogService.Logs(AdminValidate.Logs.Check(data)))
}
