package dictionary

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	DictionaryService "github.com/herman-hang/herman/app/services/dictionary"
	DictionaryValidate "github.com/herman-hang/herman/app/validates/dictionary"
)

// AddDictionary 添加数据字典
// @param *gin.Context ctx 上下文
// @return void
func AddDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	DictionaryService.Add(DictionaryValidate.Add.Check(data))
	context.Json(nil)
}

// ModifyDictionary 修改数据字典
// @param *gin.Context ctx 上下文
// @return void
func ModifyDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	DictionaryService.Modify(DictionaryValidate.Modify.Check(data))
	context.Json(nil)
}

// FindDictionary 根据ID获取数据字典详情
// @param *gin.Context ctx 上下文
// @return void
func FindDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.Find(DictionaryValidate.Find.Check(data)))
}

// RemoveDictionary 删除数据字典
// @param *gin.Context ctx 上下文
// @return void
func RemoveDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	DictionaryService.Remove(DictionaryValidate.Delete.Check(data))
	context.Json(nil)
}

// ListDictionary 数据字典列表
// @param *gin.Context ctx 上下文
// @return void
func ListDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.List(DictionaryValidate.List.Check(data)))
}

// DetailsByDictionary 数据字典KEY获取明细值
// @param *gin.Context ctx 上下文
// @return dictionary 返回数据字典和明细值
func DetailsByDictionary(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.Details(DictionaryValidate.Details.Check(data)))
}
