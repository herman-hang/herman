package admin

import (
	"github.com/herman-hang/herman/tests"
	"testing"
)

// TestLogin 测试管理员登录
// @param *testing.T t 测试对象
func TestLogin(t *testing.T) {
	tests.Call(t, []tests.TestCase{
		{
			Method:   "POST",
			Uri:      "/api/v1/admin/login",
			Params:   map[string]interface{}{"user": "admin", "password": "123456"},
			Code:     200,
			Message:  "操作成功",
			Desc:     "管理员登录",
			ShowBody: true,
		},
	})
}
