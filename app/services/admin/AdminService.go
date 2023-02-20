package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/herman/app/common"
	AdminConstant "github.com/herman/app/constants/admin"
	"github.com/herman/app/repositories"
	"github.com/herman/app/utils"
	"github.com/herman/app/validates/role"
	"github.com/herman/bootstrap/casbin"
	"gorm.io/gorm"
	"time"
)

// Login 管理员登录
// @param map[string]interface{} data 前端请求数据
// @return interface{} 返回一个token值
func Login(data map[string]interface{}) interface{} {
	// 获取管理员信息
	admin := repositories.Admin.GetAdminInfo(fmt.Sprintf("%s", data["user"]))
	// 设置上下文
	ctx := context.Background()
	// 设置Redis错误密码的key
	key := fmt.Sprintf("admin_password_error:%d", admin.Id)
	// 获取错误登录次数
	errorNumber, err := common.Redis.Get(ctx, key).Int()
	// 判断是否登录次数过多
	if err != redis.Nil && errorNumber > AdminConstant.LoginErrorLimitNumber {
		panic(AdminConstant.ErrorLoginOverload)
	}
	// 密码验证
	if !utils.ComparePasswords(admin.Password, fmt.Sprintf("%s", data["password"])) {
		common.Redis.Set(ctx, key, errorNumber+AdminConstant.Increment, time.Minute*AdminConstant.KeyValidity)
		panic(AdminConstant.PasswordError)
	}
	// 登录总数自增
	if err = repositories.Admin.Update([]uint{admin.Id}, map[string]interface{}{"login_total": admin.LoginTotal + 1}); err != nil {
		return nil
	}
	// 返回token
	return utils.GenerateToken(&utils.Claims{Uid: admin.Id, Guard: "admin"})
}

// Add 管理员添加
// @param map[string]interface{} data 前端请求数据
// @return void
func Add(data map[string]interface{}) {
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 执行添加管理员
		admin, err := repositories.Admin.Insert(data)
		if err != nil {
			return errors.New(AdminConstant.AddFail)
		}
		// 判断选择的角色是否存在
		if err := RoleKeyIsExist(data["roles"].([]role.Roles)); err != nil {
			return errors.New(AdminConstant.RoleNotExist)
		}
		// 关联角色
		if err = JoinRole(admin, data["roles"].([]role.Roles)); err != nil {
			return errors.New(AdminConstant.AddRoleFail)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// Modify 管理员修改
// @param map[string]interface{} data 前端请求数据
// @return void
func Modify(data map[string]interface{}) {
	// 过滤密码数据
	data = FilterPassword(data)
	err := common.Db.Transaction(func(tx *gorm.DB) error {
		common.Db = tx
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		id := data["id"].(uint)
		roles := data["roles"].([]role.Roles)
		delete(data, "user")
		delete(data, "roles")
		// 执行更新
		if err := repositories.Admin.Update([]uint{id}, data); err != nil {
			return errors.New(AdminConstant.UpdateFail)
		}
		// 判断选择的角色是否存在
		if err := RoleKeyIsExist(roles); err != nil {
			return errors.New(AdminConstant.RoleNotExist)
		}
		// 删除角色
		if err := repositories.AdminRole.DeleteByAdminId(id); err != nil {
			return errors.New(AdminConstant.DeleteFail)
		}
		// 关联角色
		if err := JoinRole(data, roles); err != nil {
			return errors.New(AdminConstant.AddRoleFail)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// Find 根据ID获取管理员信息
// @param map[string]interface{} data 前端请求数据
func Find(data map[string]interface{}) map[string]interface{} {
	id := data["id"].(uint)
	fields := []string{
		"id",
		"user",
		"photo",
		"name",
		"card",
		"sex",
		"age",
		"region",
		"phone",
		"email",
		"state",
		"introduction",
		"sort",
		"login_out_ip",
		"login_out_at",
		"created_at",
	}
	info, err := repositories.Admin.Find(map[string]interface{}{"id": id}, fields)
	info["roles"] = FindRole(id)
	if err != nil {
		panic(AdminConstant.GetAdminInfoFail)
	}
	return info
}
