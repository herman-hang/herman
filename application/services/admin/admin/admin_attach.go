package admin

import (
	AdminConstant "github.com/herman-hang/herman/application/constants/admin"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/application/validates/admin/role"
	"github.com/herman-hang/herman/kernel/utils"
	"gorm.io/gorm"
)

// FilterPassword 过滤密码数据
// @param data map[string]interface{} 待过滤数据
// @return map[string]interface{} 返回过滤完的数据
func FilterPassword(data map[string]interface{}) map[string]interface{} {
	if val, ok := data["password"]; ok && val != "" {
		data["password"] = utils.HashEncode(data["password"].(string))
	} else {
		delete(data, "password")
	}
	return data
}

// RoleKeyIsExist 判断角色是否存在
// []role.Roles roles 角色数据
func RoleKeyIsExist(roles []role.Roles) error {
	for _, v := range roles {
		// 判断角色Key是否存在
		if isExist, err := repositories.Role().KeyIsExist(v.Role); !isExist {
			return err
		}
	}
	return nil
}

// JoinRole 管理员关联角色
// @param map[string]interface{} admin 管理员信息
// @param []role.Roles 角色数组
// @param *gorm.DB tx 事务
// @return error 返回一个错误信息
func JoinRole(admin map[string]interface{}, roles []role.Roles, tx *gorm.DB) error {
	for _, v := range roles {
		_, err := repositories.AdminRole(tx).Insert(map[string]interface{}{
			"adminId": admin["id"],
			"roleKey": v.Role,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// FindRole 获取角色信息
// @param uint adminId 管理员ID
// @return data 返回角色数据
func FindRole(adminId uint) []map[string]interface{} {
	info, err := repositories.AdminRole().GetRoles(adminId, []string{"roles.role", "roles.name"})
	if err != nil {
		panic(AdminConstant.GetRoleFail)
	}
	return info
}
