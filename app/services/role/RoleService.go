package role

import (
	"errors"
	"fmt"
	"github.com/herman-hang/herman/app/constants"
	RoleConstant "github.com/herman-hang/herman/app/constants/role"
	"github.com/herman-hang/herman/app/repositories"
	"github.com/herman-hang/herman/app/validates/role"
	"github.com/herman-hang/herman/bootstrap/casbin"
	"github.com/herman-hang/herman/bootstrap/core"
	"gorm.io/gorm"
)

// Add 添加角色
// @param map[string]interface{} data 带处理数据
// @return void
func Add(data map[string]interface{}) {
	err := core.Db.Transaction(func(tx *gorm.DB) error {
		// casbin重新初始化
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 判断角色Key是否存在
		if isExist, _ := repositories.Role(tx).KeyIsExist(data["role"].(string)); isExist {
			return errors.New(RoleConstant.KeyExist)
		}
		roles := data["roles"]
		rules := data["rules"]
		delete(data, "roles")
		delete(data, "rules")
		// 添加角色信息
		roleInfo, err := repositories.Role(tx).Insert(data)
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
	err := core.Db.Transaction(func(tx *gorm.DB) error {
		id := data["id"].(uint)
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 判断角色是否存在
		roleInfo, _ := repositories.Role(tx).Find(map[string]interface{}{"id": id}, []string{"id", "role"})
		if len(roleInfo) == constants.LengthByZero {
			return errors.New(RoleConstant.NotExist)
		}
		roles := data["roles"]
		rules := data["rules"]
		delete(data, "roles")
		delete(data, "rules")
		// 修改角色
		if err := repositories.Role(tx).Update([]uint{id}, data); err != nil {
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

// Find 根据ID获取角色详情
// @param map[string]interface{} data 带处理数据
// @return void
func Find(data map[string]interface{}) map[string]interface{} {
	// 查询角色信息
	fields := []string{"id", "name", "role", "state", "introduction"}
	roleInfo, err := repositories.Role().Find(map[string]interface{}{"id": data["id"]}, fields)
	if err != nil {
		panic(RoleConstant.FindFail)
	}
	roles, err := core.Casbin.GetRolesForUser(roleInfo["role"].(string))
	if err != nil {
		panic(RoleConstant.FindFail)
	}
	roleInfo["roles"], err = repositories.Role().FindRoles(roles)
	if err != nil {
		panic(RoleConstant.FindFail)
	}
	// 查询菜单
	rules, err := repositories.Menu().GetAllData([]string{"path", "method", "name"})
	if err != nil {
		panic(RoleConstant.FindFail)
	}
	for i, rule := range rules {
		enforce, _ := core.Casbin.Enforce(roleInfo["role"].(string), rule["path"], rule["method"])
		fmt.Println(rule["path"].(string))
		if enforce {
			rules[i]["isPermission"] = RoleConstant.HaveAuthority
		} else {
			rules[i]["isPermission"] = RoleConstant.NotHaveAuthority
		}
	}
	roleInfo["rules"] = rules
	return roleInfo
}

// Remove 删除角色
// @param map[string]interface{} data 带处理数据
// @return void
func Remove(data map[string]interface{}) {
	err := core.Db.Transaction(func(tx *gorm.DB) error {
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 查询角色信息
		for _, id := range data["id"].([]uint) {
			roleInfo, err := repositories.Role(tx).Find(map[string]interface{}{"id": id}, []string{"id", "role"})
			if err != nil {
				return errors.New(RoleConstant.DeleteFail)
			}
			if err := repositories.Role(tx).Delete([]uint{id}); err != nil {
				return errors.New(RoleConstant.DeleteFail)
			}
			// 删除所有角色和权限
			if err := DeleteRole(roleInfo); err != nil {
				return errors.New(RoleConstant.DeleteFail)
			}
		}

		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}

// List 角色列表
// @param map[string]interface{} data 带处理数据
// @return void
func List(data map[string]interface{}) map[string]interface{} {
	// 模糊查询条件拼接
	query := fmt.Sprintf(" role like '%%%s' or name like '%%%s'", data["keywords"], data["keywords"])
	fields := []string{
		"id",
		"name",
		"role",
		"state",
		"sort",
		"introduction",
		"created_at",
	}
	// 排序
	order := "created_at desc"
	list, err := repositories.Role().List(query, fields, order, data)
	if err != nil {
		panic(RoleConstant.GetListFail)
	}
	return list
}
