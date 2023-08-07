package admin

import (
	"github.com/herman-hang/herman/application/validates"
	"github.com/herman-hang/herman/application/validates/admin/role"
)

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 管理员修改验证规则
type ModifyValidate struct {
	Id           uint         `json:"id" validate:"required,numeric" label:"管理员ID"`
	Password     string       `json:"password" validate:"omitempty,min=6,max=15" label:"密码"`
	Roles        []role.Roles `json:"roles" validate:"required" label:"选择角色"`
	PhotoId      string       `json:"photoId" validate:"omitempty" label:"头像"`
	Name         string       `json:"name" validate:"omitempty,max=20" label:"真实姓名"`
	Card         string       `json:"card" validate:"omitempty,max=20" label:"身份证号码"`
	Sex          uint8        `json:"sex" validate:"required,oneof=1 2 3" label:"性别"`
	Age          uint8        `json:"age" validate:"required,min=0,max=120" label:"年龄"`
	Region       string       `json:"region" validate:"omitempty,max=255" label:"住址"`
	Phone        string       `json:"phone" validate:"omitempty,len=11" label:"手机号码"`
	Email        string       `json:"email" validate:"omitempty,email" label:"邮箱"`
	Introduction string       `json:"introduction" validate:"omitempty" label:"简介"`
	State        uint8        `json:"state" validate:"required,oneof=1 2" label:"状态"`
	Sort         uint         `json:"sort" validate:"omitempty" label:"排序"`
}
