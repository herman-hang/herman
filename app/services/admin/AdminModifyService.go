package admin

import (
	"github.com/herman/app/repositories"
	"github.com/herman/app/utils"
	"github.com/herman/app/validates/role"
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
		if isExist, err := repositories.Role.KeyIsExist(v.Role); !isExist {
			return err
		}
	}
	return nil
}

// JoinRole 管理员关联角色
// @param map[string]interface{} admin 管理员信息
// @param []role.Roles 角色数组
// @return error 返回一个错误信息
func JoinRole(admin map[string]interface{}, roles []role.Roles) error {
	for _, v := range roles {
		_, err := repositories.AdminRole.Insert(map[string]interface{}{
			"adminId": admin["id"],
			"roleKey": v.Role,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
