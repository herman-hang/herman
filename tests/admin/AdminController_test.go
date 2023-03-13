package admin

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/bootstrap/core/test"
	"github.com/herman-hang/herman/database/seeders/admin"
	"github.com/herman-hang/herman/database/seeders/role"
	"github.com/stretchr/testify/suite"
	"testing"
)

// 管理员测试套件结构体
type AdminTestSuite struct {
	test.SuiteCase
}

var (
	AdminLoginUri = "/admin/login"  // 管理员登录URI
	AdminUri      = "/admin/admins" // 管理员URI
)

// TestLogin 测试管理员登录
// @return void
func (base *AdminTestSuite) TestLogin() {
	base.Assert([]test.Case{
		{
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "123456"},
			Code:    200,
			Message: "登录成功",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admi", "password": "123456"},
			Code:    500,
			Message: "用户名长度必须至少为5个字符",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admmin1", "password": "123456"},
			Code:    500,
			Message: "管理员不存在",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "123"},
			Code:    500,
			Message: "密码长度必须至少为6个字符",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin1111111111111", "password": "123"},
			Code:    500,
			Message: "用户名长度不能超过15个字符",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "admin", "password": "1111111111111123"},
			Code:    500,
			Message: "密码长度不能超过15个字符",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "", "password": "123456"},
			Code:    500,
			Message: "用户名为必填字段",
		}, {
			Method:  "POST",
			Uri:     base.AppPrefix + AdminLoginUri,
			Params:  map[string]interface{}{"user": "123456", "password": ""},
			Code:    500,
			Message: "密码为必填字段",
		},
	})
}

// TestAddAdmin 测试添加管理员
// @return void
func (base *AdminTestSuite) TestAddAdmin() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	adminInfo := admin.Admin()
	adminInfo["roles"] = []map[string]interface{}{
		{
			"name": roleInfo["name"].(string),
			"role": roleInfo["role"].(string),
		},
	}
	base.Assert([]test.Case{
		{
			Method:  "POST",
			Uri:     base.AppPrefix + AdminUri,
			Params:  adminInfo,
			Code:    200,
			Message: "操作成功",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + AdminUri,
			Params: map[string]interface{}{
				"user":         "sad",
				"password":     gofakeit.Password(false, false, true, false, false, 10),
				"photo":        gofakeit.ImageURL(100, 100),
				"name":         gofakeit.Name(),
				"card":         "450981200008272525",
				"sex":          gofakeit.RandomInt([]int{1, 2, 3}),
				"age":          gofakeit.Number(18, 60),
				"region":       gofakeit.Country(),
				"phone":        "18888888888",
				"email":        gofakeit.Email(),
				"introduction": gofakeit.Sentence(10),
				"state":        gofakeit.RandomInt([]int{1, 2}),
				"sort":         gofakeit.Number(1, 100),
				"roles": []map[string]interface{}{
					{
						"name": roleInfo["name"].(string),
						"role": roleInfo["role"].(string),
					},
				},
			},
			Code:    500,
			Message: "用户名长度必须至少为5个字符",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + AdminUri,
			Params: map[string]interface{}{
				"user":         "",
				"password":     gofakeit.Password(false, false, true, false, false, 10),
				"photo":        gofakeit.ImageURL(100, 100),
				"name":         gofakeit.Name(),
				"card":         "450981200008272525",
				"sex":          gofakeit.RandomInt([]int{1, 2, 3}),
				"age":          gofakeit.Number(18, 60),
				"region":       gofakeit.Country(),
				"phone":        "18888888888",
				"email":        gofakeit.Email(),
				"introduction": gofakeit.Sentence(10),
				"state":        gofakeit.RandomInt([]int{1, 2}),
				"sort":         gofakeit.Number(1, 100),
				"roles": []map[string]interface{}{
					{
						"name": roleInfo["name"].(string),
						"role": roleInfo["role"].(string),
					},
				},
			},
			Code:    500,
			Message: "用户名为必填字段",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + AdminUri,
			Params: map[string]interface{}{
				"user":         "fds456sfsa564fasf456saf",
				"password":     gofakeit.Password(false, false, true, false, false, 10),
				"photo":        gofakeit.ImageURL(100, 100),
				"name":         gofakeit.Name(),
				"card":         "450981200008272525",
				"sex":          gofakeit.RandomInt([]int{1, 2, 3}),
				"age":          gofakeit.Number(18, 60),
				"region":       gofakeit.Country(),
				"phone":        "18888888888",
				"email":        gofakeit.Email(),
				"introduction": gofakeit.Sentence(10),
				"state":        gofakeit.RandomInt([]int{1, 2}),
				"sort":         gofakeit.Number(1, 100),
				"roles": []map[string]interface{}{
					{
						"name": roleInfo["name"].(string),
						"role": roleInfo["role"].(string),
					},
				},
			},
			Code:    500,
			Message: "用户名长度不能超过15个字符",
		},
	})
}

// TestAdminTestSuite 管理员测试套件
func TestAdminTestSuite(t *testing.T) {
	suite.Run(t, &AdminTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
