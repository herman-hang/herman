package admin

import "github.com/herman-hang/herman/app/validates"

// Logs 重写验证器结构体，切记不使用引用，而是拷贝
var Logs = validates.Validates{Validate: ListValidate{}}

// ListValidate 管理员列表验证规则
type ListValidate struct {
	Page     uint   `json:"page" validate:"numeric" label:"页码"`
	PageSize uint   `json:"pageSize" validate:"numeric" label:"每页大小"`
	Keywords string `json:"keywords" validate:"omitempty,max=20" label:"每页大小"`
	Type     uint8  `json:"type" validate:"required,oneof=1 2" label:"日志类型"`
}
