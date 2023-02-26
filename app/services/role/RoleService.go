package role

import (
	"errors"
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/constants"
	RoleConstant "github.com/herman-hang/herman/app/constants/role"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/app/validates/role"
	"github.com/herman-hang/herman/bootstrap/casbin"
	"gorm.io/gorm"
)

// Add 添加角色
// @param map[string]interface{} data 带处理数据
// @return void
func Add(data map[string]interface{}) {
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		// casbin重新初始化
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 判断角色Key是否存在
		if isExist, _ := repositories.Role.KeyIsExist(data["role"].(string)); isExist {
			return errors.New(RoleConstant.KeyExist)
		}
		roles := data["roles"]
		rules := data["rules"]
		delete(data, "roles")
		delete(data, "rules")
		// 添加角色信息
		roleInfo, err := repositories.Role.Insert(data)
		if err != nil {
			return errors.New(RoleConstant.AddFail)
		}
		// 添加策略
		if err := AddPolicies(roles.([]role.Roles), rules.([]role.Rules), roleInfo); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		panic(err.Error())
	}
}

// Modify 修改角色
// @param map[string]interface{} data 带处理数据
// @return void
func Modify(data map[string]interface{}) {
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		id := data["id"].(uint)
		common.Db = tx
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 判断角色是否存在
		roleInfo, _ := repositories.Role.Find(map[string]interface{}{"id": id}, []string{"id", "role"})
		if len(roleInfo) == constants.LengthByZero {
			return errors.New(RoleConstant.NotExist)
		}
		roles := data["roles"]
		rules := data["rules"]
		delete(data, "roles")
		delete(data, "rules")
		// 修改角色
		if err := repositories.Role.Update([]uint{id}, data); err != nil {
			return errors.New(RoleConstant.ModifyFail)
		}
		// 删除所有角色和权限
		if err := DeleteRole(roleInfo); err != nil {
			return errors.New(RoleConstant.DeleteFail)
		}
		// 添加策略
		if err := AddPolicies(roles.([]role.Roles), rules.([]role.Rules), roleInfo); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}
