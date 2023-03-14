package dictionary

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/herman-hang/herman/app/constants"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/bootstrap/core"
	"github.com/herman-hang/herman/bootstrap/core/test"
	"github.com/herman-hang/herman/database/seeders/dictionary"
	"github.com/stretchr/testify/suite"
	"testing"
)

// DictionaryTestSuite 数据字典测试套件结构体
type DictionaryTestSuite struct {
	test.SuiteCase
}

var (
	DictionaryUri       = "/admin/dictionaries"         // 数据字典URI
	DictionaryDetailUri = "/admin/dictionaries/details" // 根据数据字典KEY获取明细值URI
)

// TestAddDictionary 测试添加数据字典
// @return void
func (base *DictionaryTestSuite) TestAddDictionary() {
	base.Assert([]test.Case{
		{
			Method:  "POST",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  dictionary.Dictionary(),
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestModifyDictionary 测试修改数据字典
// @return void
func (base *DictionaryTestSuite) TestModifyDictionary() {
	info, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	base.Assert([]test.Case{
		{
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   gofakeit.Name(),
				"code":   gofakeit.LetterN(5),
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestFindDictionary 测试根据ID获取数据字典详情
// @return void
func (base *DictionaryTestSuite) TestFindDictionary() {
	info, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DictionaryUri + "/" + fmt.Sprintf("%d", info["id"]),
			Params:  nil,
			Code:    200,
			Message: "操作成功",
			Fields:  []string{"id", "name", "code", "remark", "state", "createdAt"},
		},
	})
}

// TestRemoveDictionary 测试删除数据字典
// @return void
func (base *DictionaryTestSuite) TestRemoveDictionary() {
	info, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	core.Log.Debug("infoId:", info["id"])
	base.Assert([]test.Case{
		{
			Method:  "DELETE",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  map[string]interface{}{"id": []uint{info["id"].(uint)}},
			Code:    200,
			Message: "操作成功",
		},
	})
}

// TestListDictionary 测试数据字典列表
// @return void
func (base *DictionaryTestSuite) TestListDictionary() {
	_, _ = repositories.Dictionary().Insert(dictionary.Dictionary())
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
			Code:    200,
			Message: constants.Success,
			IsList:  true,
			Fields:  []string{"id", "name", "code", "remark", "state", "createdAt"},
		},
	})
}

// TestDetailsByDictionary 测试根据数据字典KEY获取明细值
// @return void
func (base *DictionaryTestSuite) TestDetailsByDictionary() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	DetailInfo := dictionary.Detail()
	DetailInfo["dictionaryId"] = dictionaryInfo["id"]
	_, _ = repositories.DictionaryDetail().Insert(DetailInfo)
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DictionaryDetailUri,
			Params:  map[string]interface{}{"keys": []string{dictionaryInfo["code"].(string)}},
			Code:    200,
			Message: "操作成功",
			Fields:  []string{"id", "name", "code", "remark", "details"},
		},
	})
}

// TestDictionaryTestSuite 数据字典测试套件
// @param *testing.T t 测试对象
// @return void
func TestDictionaryTestSuite(t *testing.T) {
	suite.Run(t, &DictionaryTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
