package repositories

import (
	"github.com/herman-hang/herman/app/constants"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/bootstrap/core"
)

// RoleRepository 角色仓储层
type RoleRepository struct {
	BaseRepository
}

// Role 实例化角色仓储层
// @return RoleRepository 返回角色仓储层
func Role() *RoleRepository {
	return &RoleRepository{BaseRepository{Model: new(models.Role)}}
}

// KeyIsExist 判断角色Key是否存在
// @param string key 角色ket
// @return bool err 返回一个bool值和一个错误信息
func (base RoleRepository) KeyIsExist(role string) (bool bool, err error) {
	result := make(map[string]interface{})
	err = core.Db.Model(&base.Model).Where("role = ?", role).Find(&result).Error
	if len(result) != constants.LengthByZero {
		return true, nil
	}
	return false, err
}

// FindRoles 查询角色信息
// @param []string roles
// @return data err 返回一个角色切片，错误信息
func (base RoleRepository) FindRoles(roles []string) (data []map[string]interface{}, err error) {
	err = core.Db.Model(&base.Model).Where("role IN ?", roles).Select([]string{"role", "name"}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
