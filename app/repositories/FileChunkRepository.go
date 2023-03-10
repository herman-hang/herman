package repositories

import "github.com/herman-hang/herman/app/models"

// FileChunkRepository 文件分片表仓储层
type FileChunkRepository struct {
	BaseRepository
}

// FileChunk 实例化文件分片表仓储层
// @return AdminRepository 返回文件分片表仓储层
func FileChunk() *FileChunkRepository {
	return &FileChunkRepository{BaseRepository{Model: new(models.FileChunk)}}
}
