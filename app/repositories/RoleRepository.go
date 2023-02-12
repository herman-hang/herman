package repositories

import (
	"github.com/herman/app/common"
	"github.com/herman/app/constants"
	"github.com/herman/app/models"
)

var Role = &RoleRepository{BaseRepository{Model: new(models.Role)}}

// RoleRepository 角色仓储层
type RoleRepository struct {
	BaseRepository
}

// KeyIsExist 判断角色Key是否存在
// @param string key 角色ket
// @return bool err 返回一个bool值和一个错误信息
func (base *BaseRepository) KeyIsExist(role string) (bool bool, err error) {
	result := make(map[string]interface{})
	err = common.Db.Model(&base.Model).Where("role = ?", role).Find(&result).Error
	if len(result) != constants.LengthByZero {
		return true, nil
	}
	return false, err
}
