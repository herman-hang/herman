package admin

import (
	"github.com/herman/app/validates"
)

var Add = &validates.Validates{Validate: new(AddValidate)}

// AddValidate 管理员添加规则
type AddValidate struct {
	User         string `json:"user" validate:"required,min=5,max=15" label:"用户名"`
	Password     string `json:"password" validate:"required,min=6,max=15" label:"密码"`
	Photo        string `json:"photo" validate:"omitempty,url,max=255" label:"头像"`
	Name         string `json:"name" validate:"omitempty,max=20" label:"真实姓名"`
	Card         string `json:"card" validate:"omitempty,max=20" label:"身份证号码"`
	Sex          uint8  `json:"sex" validate:"required,oneof=1 2 3" label:"性别"`
	Age          uint8  `json:"age" validate:"required,min=0,max=120" label:"年龄"`
	Region       string `json:"region" validate:"omitempty,max=255" label:"住址"`
	Phone        string `json:"phone" validate:"omitempty,len=11" label:"手机号码"`
	Email        string `json:"email" validate:"omitempty,email" label:"邮箱"`
	Introduction string `json:"introduction" validate:"omitempty" label:"简介"`
	State        uint8  `json:"state" validate:"required,oneof=1 2" label:"状态"`
}
