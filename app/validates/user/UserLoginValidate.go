package user

import (
	"fp-back-user/app/common"
	"github.com/mitchellh/mapstructure"
)

type Login struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func LoginValidate(data map[string]interface{}) {
	var login Login

	// map赋值给结构体
	if err := mapstructure.Decode(data, &login); err != nil {
		panic(err.Error())
	}

	err := common.Validate.Struct(login)
	if err != nil {
		panic(err)
	}
}
