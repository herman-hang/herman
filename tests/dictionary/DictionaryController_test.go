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
	"net/http"
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
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   "",
				"code":   gofakeit.Noun(),
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典名称为必填字段",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   "1111111111111111111111111111111111111111111",
				"code":   gofakeit.Noun(),
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典名称长度不能超过30个字符",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   gofakeit.Name(),
				"code":   "",
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典KEY为必填字段",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   gofakeit.Name(),
				"code":   "1111111111111111111111111111111111111111111",
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    http.StatusInternalServerError,
			Message: "数据字典KEY长度不能超过30个字符",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   gofakeit.Name(),
				"code":   gofakeit.Noun(),
				"remark": gofakeit.HackerPhrase(),
				"state":  3,
			},
			Code:    500,
			Message: "状态必须是[1 2]中的一个",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name":   gofakeit.Name(),
				"code":   gofakeit.Noun(),
				"remark": gofakeit.HackerPhrase(),
				"state":  nil,
			},
			Code:    500,
			Message: "状态为必填字段",
		}, {
			Method: "POST",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"name": gofakeit.Name(),
				"code": gofakeit.Noun(),
				"remark": `
的鹅鹅鹅饿鹅鹅鹅饿鹅鹅鹅饿呃呃呃呃呃呃呃呃呃鹅鹅鹅鹅鹅鹅饿非人防夫人夫人夫人发给热个若过过过过过热狗解耦鸡尾酒我家供热几个人欧尼酱热空降利刃洁哥连接诶个人拉开距离干劲儿梁静茹工具了看几个人老家俄国
过了容积率软连接如果了看加热就乱扔立刻就而老规矩老人家给里扔了加热炉估计软连接过了软件管理据了解热量高加了人家给乐基老规矩热量高进入了寄过来软件个老人家给老人家给老家俄国进入经二路软件管理鸡儿
感染了跟家里人科技管理热加工绿巨人立刻就GRE了几个例如接了个加热炉集隔热几个人了就放散阀`,
				"state": gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "备注长度不能超过225个字符",
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
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   "",
				"code":   gofakeit.LetterN(5),
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典名称为必填字段",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   "1111111111111111111111111111111111111111111",
				"code":   gofakeit.LetterN(5),
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典名称长度不能超过30个字符",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   gofakeit.Name(),
				"code":   "",
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "数据字典KEY为必填字段",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   gofakeit.Name(),
				"code":   "1111111111111111111111111111111111111111111",
				"remark": gofakeit.HackerPhrase(),
				"state":  gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    http.StatusInternalServerError,
			Message: "数据字典KEY长度不能超过30个字符",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   gofakeit.Name(),
				"code":   gofakeit.LetterN(5),
				"remark": gofakeit.HackerPhrase(),
				"state":  3,
			},
			Code:    500,
			Message: "状态必须是[1 2]中的一个",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":     info["id"],
				"name":   gofakeit.Name(),
				"code":   gofakeit.LetterN(5),
				"remark": gofakeit.HackerPhrase(),
				"state":  nil,
			},
			Code:    500,
			Message: "状态为必填字段",
		}, {
			Method: "PUT",
			Uri:    base.AppPrefix + DictionaryUri,
			Params: map[string]interface{}{
				"id":   info["id"],
				"name": gofakeit.Name(),
				"code": gofakeit.LetterN(5),
				"remark": `
的鹅鹅鹅饿鹅鹅鹅饿鹅鹅鹅饿呃呃呃呃呃呃呃呃呃鹅鹅鹅鹅鹅鹅饿非人防夫人夫人夫人发给热个若过过过过过热狗解耦鸡尾酒我家供热几个人欧尼酱热空降利刃洁哥连接诶个人拉开距离干劲儿梁静茹工具了看几个人老家俄国
过了容积率软连接如果了看加热就乱扔立刻就而老规矩老人家给里扔了加热炉估计软连接过了软件管理据了解热量高加了人家给乐基老规矩热量高进入了寄过来软件个老人家给老人家给老家俄国进入经二路软件管理鸡儿
感染了跟家里人科技管理热加工绿巨人立刻就GRE了几个例如接了个加热炉集隔热几个人了就放散阀`,
				"state": gofakeit.RandomInt([]int{1, 2}),
			},
			Code:    500,
			Message: "备注长度不能超过225个字符",
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
		}, {
			Method:  "DELETE",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  nil,
			Code:    500,
			Message: "数据字典ID为必填字段",
		}, {
			Method:  "DELETE",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  map[string]interface{}{"id": []uint{info["id"].(uint)}},
			Code:    500,
			Message: "删除失败",
		},
	})
}

// TestListDictionary 数据字典列表
// @return void
func (base *DictionaryTestSuite) TestListDictionary() {
	_, _ = repositories.Dictionary().Insert(dictionary.Dictionary())
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DictionaryUri,
			Params:  map[string]interface{}{"page": 1, "pageSize": 2, "keywords": ""},
			Code:    http.StatusOK,
			Message: constants.Success,
			IsList:  true,
			Fields:  []string{"id", "name", "code", "remark", "state", "createdAt"},
		},
	})
}

func (base *DictionaryTestSuite) TestDetailsDictionary() {
	dictionaryInfo, _ := repositories.Dictionary().Insert(dictionary.Dictionary())
	DetailInfo := dictionary.Detail()
	DetailInfo["dictionaryId"] = dictionaryInfo["id"]
	_, _ = repositories.DictionaryDetail().Insert(DetailInfo)
	base.Assert([]test.Case{
		{
			Method:  "GET",
			Uri:     base.AppPrefix + DictionaryDetailUri,
			Params:  map[string]interface{}{"keys": []string{dictionaryInfo["code"].(string)}},
			Code:    http.StatusOK,
			Message: constants.Success,
			Fields:  []string{"id", "name", "code", "remark", "details"},
		},
	})
}

// TestDictionaryTestSuite 字典测试套件
// @param *testing.T t 测试对象
// @return void
func TestDictionaryTestSuite(t *testing.T) {
	suite.Run(t, &DictionaryTestSuite{SuiteCase: test.SuiteCase{Guard: "admin"}})
}
