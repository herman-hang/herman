package role

import (
	"errors"
	"fmt"
	RoleConstant "github.com/herman-hang/herman/application/constants/admin/role"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/application/validates/admin/role"
	"github.com/herman-hang/herman/kernel/core"
)

// AddPolicies 角色添加策略
// @param []role.Roles roles 继承父角色
// @param []role.Rules rules 路由规则
// @param map[string]interface{} roleInfo 新增的角色信息
// @return error 返回一个错误信息
func AddPolicies(roles []role.Roles, rules []role.Rules, roleInfo map[string]interface{}) error {
	// 添加继承父角色
	if roles != nil && len(roles) != RoleConstant.ProleNotExist {
		// 判断父角色是否存在
		addExtendRoles, err := extend(roles, roleInfo)
		if err != nil {
			return err
		}
		// 添加角色
		if _, err := core.Casbin().AddGroupingPolicies(addExtendRoles); err != nil {
			return errors.New(RoleConstant.AddFail)
		}
	}
	// 添加策略规则
	if rules != nil && len(rules) != RoleConstant.RulesNotExist {
		newRole := roleInfo["role"].(string)
		for _, v := range rules {
			if _, err := core.Casbin().AddPolicy(newRole, v.Path, v.Method); err != nil {
				return errors.New(fmt.Sprintf(RoleConstant.AddRulesFail, v.Name))
			}
		}
	}
	return nil
}

// extend 判断父角色是否存在
// @param []role.Roles data 角色数据
// @param map[string]interface{} roleInfo 当前新增的角色数据
// @return roles err 返回一位数组的角色继承关系，错误信息
func extend(data []role.Roles, roleInfo map[string]interface{}) (roles [][]string, err error) {
	var newRole = roleInfo["role"].(string)
	for _, v := range data {
		prole := v.Role
		isExist, _ := repositories.Role().KeyIsExist(prole)
		if !isExist {
			return nil, errors.New(fmt.Sprintf(RoleConstant.PRoleNotExist, prole))
		}
		roles = append(roles, []string{newRole, prole})
	}

	return roles, nil
}
