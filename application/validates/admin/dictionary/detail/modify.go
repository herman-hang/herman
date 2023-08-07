package detail

import (
	"github.com/herman-hang/herman/application/validates"
)

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 明细值修改验证规则
type ModifyValidate struct {
	Id           uint   `json:"id" validate:"required,numeric" label:"明细值ID"`
	DictionaryId uint   `json:"dictionaryId" validate:"required,numeric" label:"数据字典ID"`
	Name         string `json:"name" validate:"required,max=30" label:"明细值名称"`
	Code         string `json:"code" validate:"required,max=30" label:"明细值KEY"`
	Value        string `json:"value" validate:"required,max=50" label:"明细值"`
	Remark       string `json:"remark" validate:"omitempty,max=225" label:"备注"`
	Sort         uint   `json:"sort" validate:"required,numeric" label:"排序"`
	State        uint8  `json:"state" validate:"required,oneof=1 2" label:"状态"`
}
