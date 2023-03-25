package admin

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/database/seeders/admin"
	"github.com/herman-hang/herman/database/seeders/role"
	"github.com/herman-hang/herman/kernel/core/test"
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
		},
	})
}

// TestModifyAdmin 测试修改管理员
// @return void
func (base *AdminTestSuite) TestModifyAdmin() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	adminInfo := admin.Admin()
	adminInfo["roles"] = []map[string]interface{}{
		{
			"name": roleInfo["name"].(string),
			"role": roleInfo["role"].(string),
		},
	}
	info, _ := repositories.Admin().Insert(adminInfo)
	base.Assert([]test.Case{
		{
			Method: "PUT",
			Uri:    base.AppPrefix + AdminUri,
			Params: map[string]interface{}{
				"id":           info["id"],
				"user":         gofakeit.Username(),
				"password":     gofakeit.Password(false, false, true, false, false, 10),
				"photo":        gofakeit.ImageURL(100, 100),
				"roles":        adminInfo["roles"],
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
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestDeleteAdmin 测试根据ID获取管理员详情
// @return void
func (base *AdminTestSuite) TestFindAdmin() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	adminInfo := admin.Admin()
	adminInfo["roles"] = []map[string]interface{}{
		{
			"name": roleInfo["name"].(string),
			"role": roleInfo["role"].(string),
		},
	}
	info, _ := repositories.Admin().Insert(adminInfo)
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + AdminUri + "/" + fmt.Sprintf("%d", info["id"]),
			Params:  nil,
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestGetAdminList 测试删除管理员
// @return void
func (base *AdminTestSuite) TestRemoveAdmin() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	adminInfo := admin.Admin()
	adminInfo["roles"] = []map[string]interface{}{
		{
			"name": roleInfo["name"].(string),
			"role": roleInfo["role"].(string),
		},
	}
	info, _ := repositories.Admin().Insert(adminInfo)
	base.Assert([]test.Case{
		{
			Method: "DELETE",
			Uri:    base.AppPrefix + AdminUri,
			Params: map[string]interface{}{
				"id": []uint{info["id"].(uint)},
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestGetAdminList 测试获取管理员列表
// @return void
func (base *AdminTestSuite) TestListAdmin() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	adminInfo := admin.Admin()
	adminInfo["roles"] = []map[string]interface{}{
		{
			"name": roleInfo["name"].(string),
			"role": roleInfo["role"].(string),
		},
	}
	_, _ = repositories.Admin().Insert(adminInfo)
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + AdminUri,
			Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
			Code:    200,
			Message: "操作成功",
			List:    true,
			Fields: []string{
				"id",
				"user",
				"photo",
				"sort",
				"state",
				"phone",
				"email",
				"name",
				"card",
				"introduction",
				"sex",
				"age",
				"region",
				"createdAt",
			},
		},
	})
}

// TestAdminTestSuite 管理员测试套件
// @return void
func TestAdminTestSuite(t *testing.T) {
	suite.Run(t, &AdminTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
