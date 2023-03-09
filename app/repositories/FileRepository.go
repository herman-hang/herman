package repositories

import "github.com/herman-hang/herman/app/models"

// File 实例化结构体并重写BaseRepository
var File = AdminRepository{BaseRepository{Model: new(models.File)}}

// FileRepository 管理员表仓储层
type FileRepository struct {
	BaseRepository
}
