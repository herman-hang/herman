package admin

import (
	AdminConstant "github.com/herman-hang/herman/app/constants/admin"
	"github.com/herman-hang/herman/bootstrap/core/test"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

// 管理员测试套件结构体
type AdminTestSuite struct {
	test.SuiteCase
}

var (
	AdminLoginUri = "/api/v1/admin/login" // 管理员登录URI
)

// TestLogin 测试管理员登录
// @return void
func (base *AdminTestSuite) TestLogin() {
	base.Assert([]test.Case{
		{
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "123456"},
			Code:    http.StatusOK,
			Message: AdminConstant.LoginSuccess,
			Fields:  []string{"user", "password"},
		}, {
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admi", "password": "123456"},
			Code:    http.StatusInternalServerError,
			Message: "用户名长度必须至少为5个字符",
		}, {
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admmin1", "password": "123456"},
			Code:    http.StatusInternalServerError,
			Message: "管理员不存在",
		}, {
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "123"},
			Code:    http.StatusInternalServerError,
			Message: "密码长度必须至少为6个字符",
		}, {
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin1111111111111", "password": "123"},
			Code:    http.StatusInternalServerError,
			Message: "用户名长度不能超过15个字符",
		}, {
			Method:  "POST",
			Uri:     AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "1111111111111123"},
			Code:    http.StatusInternalServerError,
			Message: "密码长度不能超过15个字符",
		},
	})
}

// TestAdminTestSuite 管理员测试套件
func TestAdminTestSuite(t *testing.T) {
	suite.Run(t, &AdminTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
