package menu

import (
	"errors"
	"fmt"
	MenuConstant "github.com/herman-hang/herman/app/constants/menu"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// Add 菜单添加
// @param map[string]interface{} data 前端请求数据
// @return void
func Add(data map[string]interface{}) {
	if _, err := repositories.Menu().Insert(data); err != nil {
		panic(MenuConstant.AddFail)
	}
}

// Modify 菜单修改
// @param map[string]interface{} data 前端请求数据
// @return void
func Modify(data map[string]interface{}) {
	if err := repositories.Menu().Update([]uint{data["id"].(uint)}, data); err != nil {
		panic(MenuConstant.ModifyFail)
	}
}

// Find 根据ID获取菜单详情
// @param map[string]interface{} data 前端请求数据
// @return map[string]interface{} 菜单信息
func Find(data map[string]interface{}) map[string]interface{} {
	info, err := repositories.Menu().Find(map[string]interface{}{"id": data["id"]},
		[]string{"id", "pid", "name", "path", "method", "sort", "created_at"},
	)
	if info["pid"] != MenuConstant.TopChild {
		topChild, err := repositories.Menu().Find(map[string]interface{}{"id": info["pid"]}, []string{"name"})
		if err != nil {
			return nil
		}
		info["pname"] = topChild["name"]
	}
	if err != nil {
		panic(MenuConstant.FindFail)
	}
	return info
}

// Remove 菜单删除
// @param map[string]interface{} data 前端请求数据
// @return void
func Remove(data map[string]interface{}) {
	err := core.Db.Transaction(func(tx *gorm.DB) error {
		ids := data["id"].([]uint)
		if err := repositories.Menu(tx).Delete(ids); err != nil {
			return errors.New(MenuConstant.DeleteFail)
		}
		// 如果存在子菜单，则全部删除
		err := repositories.Menu(tx).DeleteByMenuId(ids)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

}

// List 菜单列表
// @param map[string]interface{} data 前端请求数据
// @return map[string]interface{} 返回列表数据
func List(data map[string]interface{}) map[string]interface{} {
	// 模糊查询条件拼接
	query := fmt.Sprintf(" id like '%%%s' or name like '%%%s'", data["keywords"], data["keywords"])
	// 查询指定字段
	fields := []string{
		"id",
		"pid",
		"name",
		"path",
		"path",
		"method",
		"sort",
		"created_at",
	}
	// 排序
	order := "created_at desc"
	// 执行查询
	list, err := repositories.Menu().List(query, fields, order, data)
	if err != nil {
		panic(MenuConstant.GetListFail)
	}
	return list
}
