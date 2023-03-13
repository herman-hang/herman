package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/bootstrap/core"
	"gorm.io/gorm"
)

// FileChunkRepository 文件分片表仓储层
type FileChunkRepository struct {
	BaseRepository
}

// FileChunk 实例化文件分片表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回文件分片表仓储层
func FileChunk(tx ...*gorm.DB) *FileChunkRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &FileChunkRepository{BaseRepository{Model: new(models.FileChunk), Db: tx[0]}}
	}
	return &FileChunkRepository{BaseRepository{Model: new(models.FileChunk), Db: core.Db}}
}
