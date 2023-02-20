package admin

import "github.com/herman/app/validates"

// Find 重写验证器结构体，切记不使用引用，而是拷贝
var Find = validates.Validates{Validate: FindValidate{}}

// FindValidate 查询管理员验证规则
type FindValidate struct {
	Id uint `json:"id" validate:"required,numeric" label:"管理员ID"`
}
