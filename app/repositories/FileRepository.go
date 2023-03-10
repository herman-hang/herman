package repositories

import (
	"github.com/herman-hang/herman/app/models"
)

// FileRepository 文件表仓储层
type FileRepository struct {
	BaseRepository
}

// File 实例化文件表仓储层
// @return AdminRepository 返回文件表仓储层
func File() *FileRepository {
	return &FileRepository{BaseRepository{Model: new(models.File)}}
}
