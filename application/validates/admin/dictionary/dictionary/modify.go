package dictionary

import (
	"github.com/herman-hang/herman/application/validates"
)

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 数据字典修改验证规则
type ModifyValidate struct {
	Id     uint   `json:"id" validate:"required,numeric" label:"数据字典ID"`
	Name   string `json:"name" validate:"required,max=30" label:"数据字典名称"`
	Code   string `json:"code" validate:"required,max=30" label:"数据字典KEY"`
	Remark string `json:"remark" validate:"omitempty,max=225" label:"备注"`
	State  uint8  `json:"state" validate:"required,oneof=1 2" label:"状态"`
}
