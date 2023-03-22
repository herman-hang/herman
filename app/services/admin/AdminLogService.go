package admin

import (
	"fmt"
	AdminConstant "github.com/herman-hang/herman/app/constants/admin"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/bootstrap/core"
)

func Logs(data map[string]interface{}) map[string]interface{} {
	core.Debug(data)
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
