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
func (base AdminRoleRepository) DeleteByAdminId(id uint) error {
	err := common.Db.Where("admin_id = ?", id).Delete(&base.Model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRoles 获取管理员管理角色
// @param uint adminId 管理员ID
// @param []string fields 查询指定字段
// @return data 返回角色数据
func (base AdminRoleRepository) GetRoles(adminId uint, fields []string) (data []map[string]interface{}, err error) {
	err = common.Db.Model(&base.Model).
		Select(fields).
		Joins("JOIN roles ON roles.role = admin_role.role_key AND admin_role.admin_id = ?", adminId).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
