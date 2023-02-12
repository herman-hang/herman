package role

import (
	"fmt"
	"github.com/herman/app/common"
	RoleConstant "github.com/herman/app/constants/role"
	"github.com/herman/app/repositories"
)

// AddPolicies 角色添加策略
// @param [][]string roles 继承父角色
// @param map[string]interface{} roleInfo 新增的角色信息
// @return error 返回一个错误信息
func AddPolicies(data map[string]interface{}, roleInfo map[string]interface{}) error {
	// 添加继承父角色
	if data["roles"] != nil && len(data["roles"].([][]string)) != RoleConstant.ProleNotExist {
		// 判断角色是否存在
		exist(data["roles"].([][]string))
		// 添加角色
		if _, err := common.Casbin.AddGroupingPolicies(data["roles"].([][]string)); err != nil {
			return err
		}
	}
	// 添加策略规则
	if data["rules"] != nil && len(data["rules"].([][]string)) != RoleConstant.RulesNotExist {
		role := roleInfo["role"].(string)
		for _, v := range data["rules"].([][]string) {
			if _, err := common.Casbin.AddPolicy(role, v[0], v[1]); err != nil {
				return err
			}
		}
	}
	return nil
}

// exist 判断角色是否存在
// @param [][]string data 角色数据
// @return void
func exist(data [][]string) {
	for _, v := range data {
		isExist, _ := repositories.Role.KeyIsExist(v[1])
		if !isExist {
			panic(fmt.Sprintf(RoleConstant.NotExist, v[1]))
		}
	}
}
