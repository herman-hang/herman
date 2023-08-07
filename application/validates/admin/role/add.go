package role

import (
	"github.com/herman-hang/herman/application/validates"
)

// Add 重写验证器结构体，切记不使用引用，而是拷贝
var Add = validates.Validates{Validate: AddValidate{}}

// AddValidate 添加角色验证规则
type AddValidate struct {
	Roles        []Roles `json:"roles" validate:"omitempty" label:"父角色"`
	Name         string  `json:"name" validate:"required,max=20" label:"角色名称"`
	Role         string  `json:"role" validate:"required,max=20" label:"角色KEY"`
	State        uint8   `json:"state" validate:"required,oneof=1 2" label:"状态"`
	Introduction string  `json:"introduction" validate:"omitempty" label:"简介"`
	Rules        []Rules `json:"rules" validate:"omitempty" label:"权限"`
}

// Roles 父角色验证规则
type Roles struct {
	Name string `json:"name" validate:"max=20" label:"角色名称"`
	Role string `json:"role" validate:"max:20" label:"角色KEY"`
}

// Rules 权限规则验证
type Rules struct {
	Path   string `json:"path" validate:"max=100" label:"请求路径"`
	Method string `json:"method" validate:"max=20" label:"请求方法"`
	Name   string `json:"name" validate:"max=20" label:"菜单名称"`
}
