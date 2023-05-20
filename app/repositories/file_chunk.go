package repositories

import (
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
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

// FindChunk 查询文件分片
// @param uint fileId 文件ID
// @return []map[string]interface{} 文件分片列表
func (base FileChunkRepository) FindChunk(fileId uint) ([]map[string]interface{}, error) {
	var list []map[string]interface{}
	err := base.Db.Model(base.Model).
		Select([]string{"file_id", "chunk_number", "chunk_size", "progress", "state", "chunk_path"}).
		Where("file_id = ? AND state = ?", fileId, FileConstant.NotUploadState).
		Order("chunk_number ASC").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
