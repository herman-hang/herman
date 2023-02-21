package admin

import "github.com/herman/app/validates"

// Delete 重写验证器结构体，切记不使用引用，而是拷贝
var Delete = validates.Validates{Validate: DeleteValidate{}}

// DeleteValidate 删除管理员验证规则
type DeleteValidate struct {
	Id []uint `json:"id" validate:"required" label:"管理员ID"`
}
