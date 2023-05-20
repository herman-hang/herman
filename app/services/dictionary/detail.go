package dictionary

import (
	"fmt"
	DictionaryConstant "github.com/herman-hang/herman/app/constants/dictionary"
	"github.com/herman-hang/herman/app/repositories"
)

// AddDetail 添加数据字典明细值
// @param map[string]interface{} data 带处理数据
// @return void
func AddDetail(data map[string]interface{}) {
	if _, err := repositories.DictionaryDetail().Insert(data); err != nil {
		fmt.Println(err)
		panic(DictionaryConstant.AddFail)
	}
}

// RemoveDetail 删除数据字典明细值
// @param map[string]interface{} data 带处理数据
// @return void
func RemoveDetail(data map[string]interface{}) {
	if err := repositories.DictionaryDetail().Delete(data["id"].([]uint)); err != nil {
		panic(DictionaryConstant.DeleteFail)
	}
}

// FindDetail 根据ID获取数据字典明细值
// @param map[string]interface{} data 带处理数据
// @return void
func FindDetail(data map[string]interface{}) map[string]interface{} {
	info, err := repositories.DictionaryDetail().Find(map[string]interface{}{"id": data["id"]}, []string{
		"id", "dictionary_id", "name", "code", "value", "remark", "sort", "state", "created_at",
	})
	if len(info) == 0 {
		panic(DictionaryConstant.DetailNotExist)
	}
	if err != nil {
		panic(DictionaryConstant.FindFail)
	}
	return info
}

// ModifyDetail 修改数据字典明细值
// @param map[string]interface{} data 带处理数据
// @return void
func ModifyDetail(data map[string]interface{}) {
	if err := repositories.DictionaryDetail().Update([]uint{data["id"].(uint)}, data); err != nil {
		panic(DictionaryConstant.ModifyFail)
	}
}
