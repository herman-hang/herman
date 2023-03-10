package repositories

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/models"
)

// MenuRepository 菜单仓储层
type MenuRepository struct {
	BaseRepository
}

// Menu 实例化菜单仓储层
// @return MenuRepository 返回菜单仓储层
func Menu() *MenuRepository {
	return &MenuRepository{BaseRepository{Model: new(models.Menu)}}
}

// DeleteByMenuId 根据父菜单ID删除子菜单
// @param []uint id 管理员ID
// @return error 返回一个错误信息
func (base MenuRepository) DeleteByMenuId(id []uint) error {
	err := common.Db.Where("pid IN ?", id).Delete(&base.Model).Error
	if err != nil {
		return err
	}
	return nil
}
