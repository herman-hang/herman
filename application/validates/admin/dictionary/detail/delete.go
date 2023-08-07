package detail

import (
	"github.com/herman-hang/herman/application/validates"
)

// Delete 重写验证器结构体，切记不使用引用，而是拷贝
var Delete = validates.Validates{Validate: DeleteValidate{}}

// DeleteValidate 明细值查询验证规则
type DeleteValidate struct {
	Id []uint `json:"id" validate:"required" label:"明细值ID"`
}
