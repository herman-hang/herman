package dictionary

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/database/seeders/dictionary"
	"github.com/herman-hang/herman/kernel/core/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

var DetailUri = "/admin/dictionaries/details"

// 数据字典明细测试套件结构体
type DetailTestSuite struct {
	test.SuiteCase
}

// TestAddDetail 测试添加数据字典明细
// @return void
func (base *DetailTestSuite) TestAddDetail() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	base.Assert([]test.Case{
		{
			Method: "POST",
			Uri:    base.AppPrefix + DetailUri,
			Params: map[string]interface{}{
				"dictionaryId": dictionaryInfo["id"],
				"name":         gofakeit.Name(),
				"code":         gofakeit.LetterN(5),
				"remark":       gofakeit.HackerPhrase(),
				"value":        gofakeit.Noun(),
				"state":        gofakeit.RandomInt([]int{1, 2}),
				"sort":         gofakeit.Number(1, 100),
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestModifyDetail 测试修改数据字典明细
// @return void
func (base *DetailTestSuite) TestModifyDetail() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	detailInfo := dictionary.Detail()
	detailInfo["dictionaryId"] = dictionaryInfo["id"]
	detail, _ := repositories.DictionaryDetail().Insert(detailInfo)
	base.Assert([]test.Case{
		{
			Method: "PUT",
			Uri:    base.AppPrefix + DetailUri,
			Params: map[string]interface{}{
				"id":           detail["id"],
				"dictionaryId": dictionaryInfo["id"],
				"name":         gofakeit.Name(),
				"code":         gofakeit.LetterN(5),
				"remark":       gofakeit.HackerPhrase(),
				"value":        gofakeit.Number(1, 10),
				"state":        gofakeit.RandomInt([]int{1, 2}),
				"sort":         gofakeit.Number(1, 100),
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestRemoveDetail 测试删除数据字典明细
// @return void
func (base *DetailTestSuite) TestRemoveDetail() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	detailInfo := dictionary.Detail()
	detailInfo["dictionaryId"] = dictionaryInfo["id"]
	detail, _ := repositories.DictionaryDetail().Insert(detailInfo)
	base.Assert([]test.Case{
		{
			Method: "DELETE",
			Uri:    base.AppPrefix + DetailUri,
			Params: map[string]interface{}{
				"id": []uint{detail["id"].(uint)},
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestFindDetail 测试根据ID查找数据字典明细
// @return void
func (base *DetailTestSuite) TestFindDetail() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	detailInfo := dictionary.Detail()
	detailInfo["dictionaryId"] = dictionaryInfo["id"]
	detail, _ := repositories.DictionaryDetail().Insert(detailInfo)
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DetailUri + "/" + fmt.Sprintf("%d", detail["id"]),
			Params:  nil,
			Code:    200,
			Message: "操作成功",
			Fields:  []string{"id", "dictionaryId", "name", "code", "remark", "value", "state", "sort", "createdAt"},
		},
	})
}

// TestDetailTestSuite 数据字典明细测试套件
// @param *testing.T t 测试对象
// @return void
func TestDetailTestSuite(t *testing.T) {
	suite.Run(t, &DetailTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
