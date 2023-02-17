package role

import (
	"github.com/herman/app/common"
	RoleConstant "github.com/herman/app/constants/role"
	"github.com/herman/app/repositories"
	"github.com/herman/app/validates/role"
	"github.com/herman/bootstrap/casbin"
	"gorm.io/gorm"
)

// Add 添加角色
// @param map[string]interface{} data 带处理数据
// @return void
func Add(data map[string]interface{}) {
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 判断角色Key是否存在
		if isExist, err := repositories.Role.KeyIsExist(data["role"].(string)); isExist {
			return err
		}
		roles := data["roles"]
		rules := data["rules"]
		delete(data, "roles")
		delete(data, "rules")
		// 添加角色信息
		roleInfo, err := repositories.Role.Insert(data)
		if err != nil {
			return err
		}
		// 添加策略
		if err := AddPolicies(roles.([]role.Roles), rules.([]role.Rules), roleInfo); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		panic(RoleConstant.AddFail)
	}
}
