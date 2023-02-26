package role

import (
	"github.com/herman-hang/herman/app/common"
)

// DeleteRole 删除当前用户继承的所有角色和权限
// @param map[string]interface{} roleInfo 当前角色信息
// @return error 错误信息
func DeleteRole(roleInfo map[string]interface{}) error {
	role := roleInfo["role"].(string)
	// 删除所有继承角色
	if _, err := common.Casbin.DeleteRolesForUser(role); err != nil {
		return err
	}
	// 删除所有权限
	if _, err := common.Casbin.DeletePermissionsForUser(role); err != nil {
		return err
	}
	return nil
}
