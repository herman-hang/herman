package Menu

import "github.com/herman-hang/herman/app/validates"

// List 重写验证器结构体，切记不使用引用，而是拷贝
var List = validates.Validates{Validate: validates.ListValidate{}}
