package repositories

import (
	"github.com/herman/app/common"
	"github.com/herman/app/models"
)

// AdminRole 实例化结构体并重写BaseRepository
var AdminRole = AdminRoleRepository{BaseRepository{Model: new(models.AdminRoleModel)}}

// AdminRoleRepository 管理员角色中间表仓储层
type AdminRoleRepository struct {
	BaseRepository
}

// DeleteByAdminId 根据管理员ID删除角色
// @param uint id 管理员ID
// @return error 返回一个错误信息
func (u AdminRoleRepository) DeleteByAdminId(id uint) error {
	err := common.Db.Where("admin_id = ?", id).Delete(&u.Model).Error
	if err != nil {
		return err
	}
	return nil
}
