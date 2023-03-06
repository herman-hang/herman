package Menu

import "github.com/herman-hang/herman/app/validates"

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 菜单添加验证规则
type ModifyValidate struct {
	Id     uint   `json:"id" validate:"required,numeric" label:"菜单ID"`
	Pid    uint   `json:"pid" validate:"omitempty,numeric" label:"父菜单ID"`
	Name   string `json:"name" validate:"required,max=20" label:"菜单名称"`
	Path   string `json:"path" validate:"required,max=100" label:"菜单路由"`
	Method string `json:"method" validate:"required,max=20" label:"请求方法"`
	Sort   uint   `json:"sort" validate:"omitempty,numeric" label:"排序"`
}
