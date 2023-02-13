package role

import (
	"github.com/herman/app/constants"
	"github.com/herman/app/utils"
	"github.com/herman/app/validates"
	"github.com/mitchellh/mapstructure"
)

// AddValidate 添加角色验证器
type AddValidate struct {
	Roles        [][]string `json:"roles" validate:"omitempty" label:"父角色KEY"`
	Name         string     `json:"name" validate:"required,max=20" label:"角色名称"`
	Role         string     `json:"role" validate:"required,max=20" label:"角色KEY"`
	State        uint8      `json:"state" validate:"required,oneof=1 2" label:"状态"`
	Introduction string     `json:"introduction" validate:"omitempty" label:"简介"`
	Rules        [][]string `json:"rules" validate:"omitempty" label:"权限"`
}

// Add 添加角色验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func Add(data map[string]interface{}) (toMap map[string]interface{}) {
	var add AddValidate
	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &add); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(add); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&add, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}
