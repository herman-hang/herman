package dictionary

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	DictionaryService "github.com/herman-hang/herman/application/services/admin/dictionary"
	"github.com/herman-hang/herman/application/validates/admin/dictionary/dictionary"
)

// AddDictionary 添加数据字典
// @param *gin.Context ctx 上下文
// @return void
func AddDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.AddDictionary(dictionary.Add.Check(data))
	context.Json(nil)
}

// ModifyDictionary 修改数据字典
// @param *gin.Context ctx 上下文
// @return void
func ModifyDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.ModifyDictionary(dictionary.Modify.Check(data))
	context.Json(nil)
}

// FindDictionary 根据ID获取数据字典详情
// @param *gin.Context ctx 上下文
// @return void
func FindDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.FindDictionary(dictionary.Find.Check(data)))
}

// RemoveDictionary 删除数据字典
// @param *gin.Context ctx 上下文
// @return void
func RemoveDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.RemoveDictionary(dictionary.Delete.Check(data))
	context.Json(nil)
}

// ListDictionary 数据字典列表
// @param *gin.Context ctx 上下文
// @return void
func ListDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.ListDictionary(dictionary.List.Check(data)))
}

// DetailsByDictionary 数据字典KEY获取明细值
// @param *gin.Context ctx 上下文
// @return dictionary 返回数据字典和明细值
func DetailsByDictionary(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.DetailsDictionary(dictionary.Details.Check(data)))
}
