package system

import "github.com/herman-hang/herman/app/validates"

// Modify 重写验证器结构体，切记不使用引用，而是拷贝
var Modify = validates.Validates{Validate: ModifyValidate{}}

// ModifyValidate 修改系统设置验证规则
type ModifyValidate struct {
	Name        string `json:"name" validate:"required,max=20" label:"网站名称"`
	Title       string `json:"title" validate:"omitempty,max=70" label:"网站标题"`
	Description string `json:"description" validate:"omitempty,max=200" label:"网站描述"`
	Keywords    string `json:"keywords" validate:"omitempty,max=100" label:"关键词"`
	LogoFileId  uint   `json:"logoFileId" validate:"omitempty" label:"LOGO文件ID"`
	IcoFileId   uint   `json:"icoFileId" validate:"omitempty" label:"LOGO文件ID"`
	Record      string `json:"record" validate:"omitempty,max=20" label:"备案号"`
	Copyright   string `json:"copyright" validate:"omitempty,max=100" label:"底部版权声明"`
	IsWebsite   uint8  `json:"isWebsite" validate:"omitempty,oneof=1 2" label:"网站开关"`
	Email       string `json:"email" validate:"omitempty,max=32" label:"邮箱"`
	Telephone   string `json:"telephone" validate:"omitempty,max=80" label:"电话"`
	Address     string `json:"address" validate:"omitempty,max=255" label:"地址"`
}
