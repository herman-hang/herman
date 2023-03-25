package role

import "github.com/herman-hang/herman/app/validates"

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 修改角色验证规则
type ModifyValidate struct {
	Id           uint    `json:"id" validate:"required,numeric" label:"角色ID"`
	Roles        []Roles `json:"roles" validate:"omitempty" label:"父角色"`
	Name         string  `json:"name" validate:"required,max=20" label:"角色名称"`
	State        uint8   `json:"state" validate:"required,oneof=1 2" label:"状态"`
	Introduction string  `json:"introduction" validate:"omitempty" label:"简介"`
	Rules        []Rules `json:"rules" validate:"omitempty" label:"权限"`
}
