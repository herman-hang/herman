package role

import (
	"fmt"
	"github.com/herman/app/common"
	RoleConstant "github.com/herman/app/constants/role"
	"github.com/herman/app/repositories"
	"github.com/herman/app/validates/role"
)

// AddPolicies 角色添加策略
// @param []role.Roles roles 继承父角色
// @param []role.Rules rules 路由规则
// @param map[string]interface{} roleInfo 新增的角色信息
// @return error 返回一个错误信息
func AddPolicies(roles []role.Roles, rules []role.Rules, roleInfo map[string]interface{}) error {
	// 添加继承父角色
	if roles != nil && len(roles) != RoleConstant.ProleNotExist {
		extendRoles := extend(roles, roleInfo)
		// 添加角色
		if _, err := common.Casbin.AddGroupingPolicies(extendRoles); err != nil {
			return err
		}
	}
	// 添加策略规则
	if rules != nil && len(rules) != RoleConstant.RulesNotExist {
		newRole := roleInfo["role"].(string)
		for _, v := range rules {
			if _, err := common.Casbin.AddPolicy(newRole, v.Path, v.Method); err != nil {
				return err
			}
		}
	}
	return nil
}

// extend 判断角色是否存在
// @param []role.Roles data 角色数据
// @param map[string]interface{} roleInfo 当前新增的角色数据
// @return roles 返回一位数组的角色继承关系
func extend(data []role.Roles, roleInfo map[string]interface{}) (roles [][]string) {
	var newRole = roleInfo["role"].(string)
	for _, v := range data {
		prole := v.Role
		isExist, _ := repositories.Role.KeyIsExist(prole)
		if !isExist {
			panic(fmt.Sprintf(RoleConstant.NotExist, prole))
		}
		roles = append(roles, []string{newRole, prole})
	}

	return roles
}
