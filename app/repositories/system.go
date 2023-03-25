package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// SystemRepository 系统设置表仓储层
type SystemRepository struct {
	BaseRepository
}

// System 实例化系统设置表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回系统设置表仓储层
func System(tx ...*gorm.DB) *SystemRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &SystemRepository{BaseRepository{Model: new(models.System), Db: tx[0]}}
	}

	return &SystemRepository{BaseRepository{Model: new(models.System), Db: core.Db}}
}
