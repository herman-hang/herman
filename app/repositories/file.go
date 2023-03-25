package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// FileRepository 文件表仓储层
type FileRepository struct {
	BaseRepository
}

// File 实例化文件表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回文件表仓储层
func File(tx ...*gorm.DB) *FileRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &FileRepository{BaseRepository{Model: new(models.File), Db: tx[0]}}
	}
	return &FileRepository{BaseRepository{Model: new(models.File), Db: core.Db}}
}
