package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	AdminLogService "github.com/herman-hang/herman/app/services/admin"
	AdminValidate "github.com/herman-hang/herman/app/validates/admin"
)

// LogList 管理员日志列表
// @param *gin.Context ctx 上下文
// @return void
func LogList(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(AdminLogService.Logs(AdminValidate.Logs.Check(data)))
}
