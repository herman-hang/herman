package dictionary

import "github.com/herman-hang/herman/application/validates"

// Details 重写验证器结构体，切记不使用引用，而是拷贝
var Details = validates.Validates{Validate: DetailsValidate{}}

// DetailsValidate 数据字典Key查询明细值验证规则
type DetailsValidate struct {
	Keys []string `json:"keys" validate:"required" label:"数据字典KEY"`
}
