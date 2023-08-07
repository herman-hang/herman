package system

import (
	"github.com/herman-hang/herman/database/seeders/system"
	"github.com/herman-hang/herman/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// 系统设置测试套件结构体
type SystemTestSuite struct {
	tests.SuiteCase
}

var SystemUri = "/admin/system" // 系统设置URI

// TestFind 测试获取系统设置信息
// @return void
func (base *SystemTestSuite) TestFind() {
	base.Assert([]tests.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + SystemUri,
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestModify 测试修改系统设置信息
func (base *SystemTestSuite) TestModify() {
	base.Assert([]tests.Case{
		{
			Method:  "PUT",
			Uri:     base.AppPrefix + SystemUri,
			Params:  system.System(),
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestAdminTestSuite 管理员测试套件
// @return void
func TestSystemTestSuite(t *testing.T) {
	suite.Run(t, &SystemTestSuite{SuiteCase: tests.SuiteCase{Guard: "admin"}})
}
