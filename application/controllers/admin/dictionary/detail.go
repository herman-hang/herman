package dictionary

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	DictionaryService "github.com/herman-hang/herman/application/services/admin/dictionary"
	DetailValidate "github.com/herman-hang/herman/application/validates/admin/dictionary/detail"
)

// AddDetail 添加数据字典明细值
// @param *gin.Context ctx 上下文
// @return void
func AddDetail(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.AddDetail(DetailValidate.Add.Check(data))
	context.Json(nil)
}

// RemoveDetail 删除数据字典明细值
// @param *gin.Context ctx 上下文
// @return void
func RemoveDetail(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.RemoveDetail(DetailValidate.Delete.Check(data))
	context.Json(nil)
}

// FindDetail 根据ID获取数据字典明细值
// @param *gin.Context ctx 上下文
// @return void
func FindDetail(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(DictionaryService.FindDetail(DetailValidate.Find.Check(data)))
}

// ModifyDetail 修改数据字典明细值
//// @param *gin.Context ctx 上下文
//// @return void
func ModifyDetail(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	DictionaryService.ModifyDetail(DetailValidate.Modify.Check(data))
	context.Json(nil)
}
