package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	AdminConstant "github.com/herman-hang/herman/app/constants/admin"
	"github.com/herman-hang/herman/app/constants/middleware"
	"github.com/herman-hang/herman/app/repositories"
	"net/http"
)

// Logs 管理员日志列表
// @param map[string]interface{} data 待处理数据
// @return map[string]interface{} 返回数据
func Logs(data map[string]interface{}) map[string]interface{} {
	// 模糊查询条件拼接
	query := fmt.Sprintf(" type = %d and admin_id like '%%%s' or id like '%%%s' or path like '%%%s'",
		data["type"], data["keywords"], data["keywords"], data["keywords"])
	// 查询指定字段
	fields := []string{
		"id",
		"admin_id",
		"ip",
		"path",
		"method",
		"remark",
		"code",
		"state",
		"created_at",
	}
	// 排序
	order := "created_at desc"
	// 执行查询
	list, err := repositories.AdminLog().List(query, fields, order, data)
	if err != nil {
		panic(AdminConstant.GetAdminListFail)
	}
	return list
}

// LogWriter 管理员日志记录
// @param string user 管理员用户名
// @param uint16 code 状态码
// @param string remark 备注
// @param *gin.Context ctx 上下文
// @return void
func LogWriter(user string, code uint16, remark string, ctx *gin.Context) {
	var state int
	admin := repositories.Admin().GetAdminInfo(user)
	if code != http.StatusOK {
		state = 1
	} else {
		state = 2
	}
	_, _ = repositories.AdminLog().Insert(map[string]interface{}{
		"type":    middleware.LoginType,
		"adminId": admin.Id,
		"ip":      ctx.ClientIP(),
		"path":    ctx.Request.URL.Path,
		"method":  ctx.Request.Method,
		"code":    code,
		"remark":  remark,
		"state":   state,
	})
}
