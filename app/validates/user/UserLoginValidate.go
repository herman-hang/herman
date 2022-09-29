package user

import (
	"fp-back-user/app/validates"
	"github.com/mitchellh/mapstructure"
)

type LoginValidate struct {
	User     string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=15" label:"密码"`
}

func Login(data map[string]interface{}) map[string]interface{} {
	var login LoginValidate

	// map赋值给结构体
	if err := mapstructure.Decode(data, &login); err != nil {
		panic(err.Error())
	}

	if err := validates.Validate(login); err != nil {
		panic(err.Error())
	}

	return data
}
