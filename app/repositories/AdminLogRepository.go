package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// AdminLogRepository 管理员日志表仓储层
type AdminLogRepository struct {
	BaseRepository
}

// AdminLog 实例化管理员日志表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回管理员日志表仓储层
func AdminLog(tx ...*gorm.DB) *AdminLogRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &AdminLogRepository{BaseRepository{Model: new(models.AdminLog), Db: tx[0]}}
	}

	return &AdminLogRepository{BaseRepository{Model: new(models.AdminLog), Db: core.Db}}
}
