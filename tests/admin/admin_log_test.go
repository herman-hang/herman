package admin

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// 管理员日志测试套件结构体
type AdminLogTestSuite struct {
	tests.SuiteCase
}

var AdminLogUri = "/admin/admin/logs" // 管理员日志列表URI

// 测试管理员日志列表
// @return void
func (base *AdminLogTestSuite) TestLogList() {
	_, _ = repositories.AdminLog().Insert(map[string]interface{}{
		"type":    gofakeit.RandomInt([]int{1, 2}),
		"adminId": 1,
		"ip":      gofakeit.IPv4Address(),
		"path":    gofakeit.URL(),
		"method":  gofakeit.RandomString([]string{"POST", "GET", "PUT", "DELETE"}),
		"remark":  gofakeit.Sentence(1),
	})
	base.Assert([]tests.Case{
		{
			Method: "GET",
			Uri:    base.AppPrefix + AdminLogUri,
			Params: map[string]interface{}{
				"type":     gofakeit.RandomInt([]int{1, 2}),
				"page":     1,
				"pageSize": 2,
				"keywords": "",
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestAdminLogTestSuite 管理员测试套件
// @return void
func TestAdminLogTestSuite(t *testing.T) {
	suite.Run(t, &AdminLogTestSuite{SuiteCase: tests.SuiteCase{Guard: "admin"}})
}
