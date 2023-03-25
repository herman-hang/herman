package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	MenuService "github.com/herman-hang/herman/app/services/menu"
	MenuValidate "github.com/herman-hang/herman/app/validates/menu"
)

// AddMenu 菜单添加
// @param *gin.Context ctx 上下文
// @return void
func AddMenu(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	MenuService.Add(MenuValidate.Add.Check(data))
	context.Json(nil)
}

// ModifyMenu 菜单修改
// @param *gin.Context ctx 上下文
// @return void
func ModifyMenu(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	MenuService.Modify(MenuValidate.Modify.Check(data))
	context.Json(nil)
}

// FindMenu 根据ID获取菜单详情
// @param *gin.Context ctx 上下文
// @return void
func FindMenu(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(MenuService.Find(MenuValidate.Find.Check(data)))
}

// RemoveMenu 菜单删除
// @param *gin.Context ctx 上下文
// @return void
func RemoveMenu(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	MenuService.Remove(MenuValidate.Delete.Check(data))
	context.Json(nil)
}

// ListMenu 菜单列表
// @param *gin.Context ctx 上下文
// @return void
func ListMenu(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	context.Json(MenuService.List(MenuValidate.List.Check(data)))
}
