package menu

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/database/seeders/menu"
	"github.com/herman-hang/herman/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestSuite 菜单测试套件
type MenuTestSuite struct {
	tests.SuiteCase
}

var MenuUri = "/admin/menus"

// TestAddMenu 测试添加菜单
// @return void
func (base *MenuTestSuite) TestAddMenu() {
	base.Assert([]tests.Case{
		{
			Method:  "POST",
			Uri:     base.AppPrefix + MenuUri,
			Params:  menu.Menu(),
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestModifyMenu 测试修改菜单
// @return void
func (base *MenuTestSuite) TestModifyMenu() {
	menuInfo, _ := repositories.Menu().Insert(menu.Menu())
	base.Assert([]tests.Case{
		{
			Method: "PUT",
			Uri:    base.AppPrefix + MenuUri,
			Params: map[string]interface{}{
				"id":   menuInfo["id"],
				"pid":  0,
				"name": gofakeit.Name(),
				"path": gofakeit.URL(),
				"method": gofakeit.RandomString([]string{
					"GET", "POST", "PUT", "DELETE",
				}),
				"sort": gofakeit.Number(0, 100),
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestDeleteMenu 测试根据ID获取菜单详情
// @return void
func (base *MenuTestSuite) TestFindMenu() {
	menuInfo, _ := repositories.Menu().Insert(menu.Menu())
	base.Assert([]tests.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + MenuUri + "/" + fmt.Sprintf("%d", menuInfo["id"]),
			Params:  nil,
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestRemoveMenu 测试删除菜单
// @return void
func (base *MenuTestSuite) TestRemoveMenu() {
	menuInfo, _ := repositories.Menu().Insert(menu.Menu())
	base.Assert([]tests.Case{
		{
			Method: "DELETE",
			Uri:    base.AppPrefix + MenuUri,
			Params: map[string]interface{}{
				"id": []uint{menuInfo["id"].(uint)},
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestListMenu 测试获取菜单列表
// @return void
func (base *MenuTestSuite) TestListMenu() {
	_, _ = repositories.Menu().Insert(menu.Menu())
	base.Assert([]tests.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + MenuUri,
			Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
			Code:    200,
			Message: "操作成功",
			List:    true,
			Fields: []string{
				"id",
				"pid",
				"name",
				"path",
				"method",
				"sort",
				"createdAt",
			},
		},
	})
}

// TestAdminTestSuite 角色测试套件
// @return void
func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, &MenuTestSuite{SuiteCase: tests.SuiteCase{Guard: "admin"}})
}
