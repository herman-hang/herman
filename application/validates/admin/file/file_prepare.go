package file

import (
	"github.com/herman-hang/herman/application/validates"
)

// Prepare 重写验证器结构体，切记不使用引用，而是拷贝
var Prepare = validates.Validates{Validate: PrepareValidate{}}

// PrepareValidate 管理员添加验证规则
type PrepareValidate struct {
	FileName string `json:"fileName" validate:"required,max=255" label:"文件名称"`
	FileExt  string `json:"fileExt" validate:"required,max=30" label:"文件扩展名"`
	FileType string `json:"fileType" validate:"required,max=30" label:"文件MIME类型"`
	Hash     string `json:"hash" validate:"required,md5" label:"文件校验和"`
	FileSize uint64 `json:"fileSize" validate:"required,number" label:"文件大小"`
}
