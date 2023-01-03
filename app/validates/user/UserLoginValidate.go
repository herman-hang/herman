package user

import (
	"github.com/fp/fp-gin-framework/app/constants"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/fp/fp-gin-framework/app/validates"
	"github.com/mitchellh/mapstructure"
)

type LoginValidate struct {
	User     string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=15" label:"密码"`
}

// Login 登录验证器
// @param map[string]interface{} data 待验证数据
// @return toMap 返回验证通过的数据
func Login(data map[string]interface{}) (toMap map[string]interface{}) {
	var login LoginValidate

	// map赋值给结构体
	if err := mapstructure.WeakDecode(data, &login); err != nil {
		panic(constants.MapToStruct)
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	toMap, err := utils.ToMap(&login, "json")
	if err != nil {
		panic(constants.StructToMap)
	}

	return toMap
}
