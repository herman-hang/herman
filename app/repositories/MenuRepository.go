package repositories

import (
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/bootstrap/core"
	"gorm.io/gorm"
)

// MenuRepository 菜单仓储层
type MenuRepository struct {
	BaseRepository
}

// Menu 实例化菜单仓储层
// @param *gorm.DB tx 事务
// @return MenuRepository 返回菜单仓储层
func Menu(tx ...*gorm.DB) *MenuRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &MenuRepository{BaseRepository{Model: new(models.Menu), Db: tx[0]}}
	}
	return &MenuRepository{BaseRepository{Model: new(models.Menu), Db: core.Db}}
}

// DeleteByMenuId 根据父菜单ID删除子菜单
// @param []uint id 管理员ID
// @return error 返回一个错误信息
func (base MenuRepository) DeleteByMenuId(id []uint) error {
	err := core.Db.Where("pid IN ?", id).Delete(&base.Model).Error
	if err != nil {
		return err
	}
	return nil
}
