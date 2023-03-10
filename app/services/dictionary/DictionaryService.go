package dictionary

import (
	"errors"
	"fmt"
	"github.com/herman-hang/herman/app/common"
	DictionaryConstant "github.com/herman-hang/herman/app/constants/dictionary"
	"github.com/herman-hang/herman/app/repositories"
	"gorm.io/gorm"
)

// AddDictionary 添加数据字典
// @param map[string]interface{} data 带处理数据
// @return void
func AddDictionary(data map[string]interface{}) {
	if _, err := repositories.Dictionary().Insert(data); err != nil {
		panic(DictionaryConstant.AddFail)
	}
}

// ModifyDictionary 修改数据字典
// @param map[string]interface{} data 带处理数据
// @return void
func ModifyDictionary(data map[string]interface{}) {
	if err := repositories.Dictionary().Update([]uint{data["id"].(uint)}, data); err != nil {
		panic(DictionaryConstant.ModifyFail)
	}
}

// FindDictionary 根据ID获取数据字典详情
// @param map[string]interface{} data 带处理数据
// @return map[string]interface{} 数据字典信息
func FindDictionary(data map[string]interface{}) map[string]interface{} {
	info, err := repositories.Dictionary().Find(map[string]interface{}{"id": data["id"]}, []string{
		"id", "name", "code", "remark", "state", "created_at",
	})
	if err != nil {
		panic(DictionaryConstant.FindFail)
	}
	return info
}

// RemoveDictionary 删除数据字典
// @param map[string]interface{} data 带处理数据
// @return void
func RemoveDictionary(data map[string]interface{}) {
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		ids := data["id"].([]uint)
		if err := repositories.Dictionary().Delete(ids); err != nil {
			return errors.New(DictionaryConstant.DeleteFail)
		}
		// 删除数据字典下的明细值
		_ = repositories.Dictionary().DeleteByDictionaryId(ids)
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}

// ListDictionary 数据字典列表
// @param map[string]interface{} data 带处理数据
// @return map[string]interface{} 返回列表数据
func ListDictionary(data map[string]interface{}) map[string]interface{} {
	// 模糊查询条件拼接
	query := fmt.Sprintf(" id like '%%%s' or name like '%%%s' or code like '%%%s'", data["keywords"], data["keywords"], data["keywords"])
	// 查询指定字段
	fields := []string{
		"id",
		"name",
		"code",
		"remark",
		"state",
		"created_at",
	}
	// 排序
	order := "created_at desc"
	list, err := repositories.Dictionary().GetList(query, fields, order, data)
	if err != nil {
		panic(DictionaryConstant.GetListFail)
	}
	return list
}

// DetailsDictionary 数据字典KEY获取明细值
// @param map[string]interface{} data 带处理数据
// @return dictionary 返回数据字典和明细值
func DetailsDictionary(data map[string]interface{}) (dictionary []map[string]interface{}) {
	var list []map[string]interface{}
	keys := data["keys"].([]string)
	if len(keys) > 0 {
		for _, key := range keys {
			dictionary, err := repositories.Dictionary().Find(map[string]interface{}{"code": key}, []string{
				"id", "name", "code", "remark",
			})
			if err != nil {
				panic(DictionaryConstant.FindFail)
			}
			dictionaryDetail, err := repositories.DictionaryDetail().FindByCode(map[string]interface{}{"dictionary_id": dictionary["id"]}, []string{
				"id", "name", "value", "remark", "sort",
			})
			if err != nil {
				panic(DictionaryConstant.FindFail)
			}
			dictionary["details"] = dictionaryDetail
			list = append(list, dictionary)
		}
		return list
	}
	return nil
}
