package role

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/database/seeders/role"
	"github.com/herman-hang/herman/kernel/core/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// RoleTestSuite 角色测试套件结构体
type RoleTestSuite struct {
	test.SuiteCase
}

var RoleUri = "/admin/roles"

// TestAddRole 测试添加角色
// @return void
func (base *RoleTestSuite) TestAddRole() {
	base.Assert([]test.Case{
		{
			Method:  "POST",
			Uri:     base.AppPrefix + RoleUri,
			Params:  role.Role(),
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestModifyRole 测试修改角色
// @return void
func (base *RoleTestSuite) TestModifyRole() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	base.Assert([]test.Case{
		{
			Method: "PUT",
			Uri:    base.AppPrefix + RoleUri,
			Params: map[string]interface{}{
				"id":    roleInfo["id"],
				"roles": nil,
				"name":  gofakeit.Name(),
				"role":  gofakeit.Username(),
				"state": gofakeit.RandomInt([]int{1, 2}),
				"rules": []map[string]interface{}{
					{
						"path":   gofakeit.PhoneFormatted(),
						"method": gofakeit.RandomString([]string{"POST", "GET", "PUT", "DELETE"}),
						"name":   gofakeit.Name(),
					},
				},
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestFindRole 测试根据ID获取角色详情
// @return void
func (base *RoleTestSuite) TestFindRole() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + RoleUri + "/" + fmt.Sprintf("%d", roleInfo["id"]),
			Params:  nil,
			Code:    200,
			Message: "操作成功",
			Fields: []string{
				"id",
				"name",
				"role",
				"state",
				"introduction",
			},
		},
	})
}

// TestGetRoleList 测试获取角色列表
// @return void
func (base *RoleTestSuite) TestRemoveRole() {
	roleInfo, _ := repositories.Role().Insert(role.Role())
	base.Assert([]test.Case{
		{
			Method: "DELETE",
			Uri:    base.AppPrefix + RoleUri,
			Params: map[string]interface{}{
				"id": []uint{roleInfo["id"].(uint)},
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestGetRoleList 测试获取角色列表
// @return void
func (base *RoleTestSuite) TestListRole() {
	_, _ = repositories.Role().Insert(role.Role())
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + RoleUri,
			Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
			Code:    200,
			Message: "操作成功",
			IsList:  true,
			Fields: []string{
				"id",
				"name",
				"role",
				"sort",
				"state",
				"introduction",
				"createdAt",
			},
		},
	})
}

// TestAdminTestSuite 角色测试套件
// @return void
func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, &RoleTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
