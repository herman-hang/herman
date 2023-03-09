package repositories

import "github.com/herman-hang/herman/app/models"

// FileChunk 实例化结构体并重写BaseRepository
var FileChunk = AdminRepository{BaseRepository{Model: new(models.FileChunk)}}

// FileChunkRepository 管理员表仓储层
type FileChunkRepository struct {
	BaseRepository
}
