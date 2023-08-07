package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	AdminConstant "github.com/herman-hang/herman/application/constants/admin"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/application/validates/admin/role"
	"github.com/herman-hang/herman/kernel/casbin"
	"github.com/herman-hang/herman/kernel/core"
	utils2 "github.com/herman-hang/herman/kernel/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Login 管理员登录
// @param map[string]interface{} data 前端请求数据
// @param *gin.Context c 上下文
// @return interface{} 返回一个token值
func Login(data map[string]interface{}, c *gin.Context) interface{} {
	// 获取管理员信息
	admin := repositories.Admin().GetAdminInfo(fmt.Sprintf("%s", data["user"]))
	if len(admin.User) == AdminConstant.NotExist {
		panic(AdminConstant.UserNotExist)
	}
	// 设置上下文
	ctx := context.Background()
	// 设置Redis错误密码的key
	key := fmt.Sprintf("admin_password_error:%d", admin.Id)
	// 获取错误登录次数
	errorNumber, err := core.Redis().Get(ctx, key).Int()
	// 判断是否登录次数过多
	if err != redis.Nil && errorNumber > AdminConstant.LoginErrorLimitNumber {
		LogWriter(data["user"].(string), http.StatusInternalServerError, AdminConstant.ErrorLoginOverload, c)
		panic(AdminConstant.ErrorLoginOverload)
	}
	// 密码验证
	if !utils2.ComparePasswords(admin.Password, fmt.Sprintf("%s", data["password"])) {
		core.Redis().Set(ctx, key, errorNumber+AdminConstant.Increment, time.Minute*AdminConstant.KeyValidity)
		LogWriter(data["user"].(string), http.StatusInternalServerError, AdminConstant.PasswordError, c)
		panic(AdminConstant.PasswordError)
	}
	// 登录总数自增
	if err = repositories.Admin().Update([]uint{admin.Id}, map[string]interface{}{"loginTotal": admin.LoginTotal + 1}); err != nil {
		return nil
	}
	LogWriter(data["user"].(string), http.StatusOK, AdminConstant.LoginSuccess, c)
	// 返回token
	return utils2.GenerateToken(&utils2.Claims{Uid: admin.Id, Guard: "admin"})
}

// Add 管理员添加
// @param map[string]interface{} data 前端请求数据
// @return void
func Add(data map[string]interface{}) {
	err := core.Db().Transaction(func(tx *gorm.DB) error {
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		// 执行添加管理员
		adminRepository := repositories.Admin(tx)
		admin, err := adminRepository.Insert(data)
		if err != nil {
			return errors.New(AdminConstant.AddFail)
		}
		// 判断选择的角色是否存在
		if err := RoleKeyIsExist(data["roles"].([]role.Roles)); err != nil {
			return errors.New(AdminConstant.RoleNotExist)
		}
		// 关联角色
		if err = JoinRole(admin, data["roles"].([]role.Roles), tx); err != nil {
			return errors.New(AdminConstant.AddRoleFail)
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}

// Modify 管理员修改
// @param map[string]interface{} data 前端请求数据
// @return void
func Modify(data map[string]interface{}) {
	// 过滤密码数据
	data = FilterPassword(data)
	err := core.Db().Transaction(func(tx *gorm.DB) error {
		_, _ = casbin.InitEnforcer(casbin.GetAdminPolicy(), tx)
		id := data["id"].(uint)
		roles := data["roles"].([]role.Roles)
		delete(data, "user")
		delete(data, "roles")
		// 执行更新
		if err := repositories.Admin(tx).Update([]uint{id}, data); err != nil {
			return errors.New(AdminConstant.UpdateFail)
		}
		// 判断选择的角色是否存在
		if err := RoleKeyIsExist(roles); err != nil {
			return errors.New(AdminConstant.RoleNotExist)
		}
		// 删除角色
		if err := repositories.AdminRole(tx).DeleteByAdminId(id); err != nil {
			return errors.New(AdminConstant.DeleteFail)
		}
		// 关联角色
		if err := JoinRole(data, roles, tx); err != nil {
			return errors.New(AdminConstant.AddRoleFail)
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}

// Find 根据ID获取管理员信息
// @param map[string]interface{} data 前端请求数据
// @return void
func Find(data map[string]interface{}) map[string]interface{} {
	id := data["id"].(uint)
	fields := []string{
		"id",
		"user",
		"photo_id",
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
	info, err := repositories.Admin().Find(map[string]interface{}{"id": id}, fields)
	if len(info) == 0 {
		panic(AdminConstant.UserNotExist)
	}
	if err != nil {
		panic(AdminConstant.GetAdminInfoFail)
	}
	info["roles"] = FindRole(id)
	return info
}

// Remove 删除管理员
// @param map[string]interface{} data 前端请求数据
// @return void
func Remove(data map[string]interface{}) {
	if err := repositories.Admin().Delete(data["id"].([]uint)); err != nil {
		panic(AdminConstant.DeleteAdminFail)
	}
}

// List 管理员列表
// @param map[string]interface{} data 前端请求数据
// @return map[string]interface{} 返回列表数据
func List(data map[string]interface{}) map[string]interface{} {
	// 模糊查询条件拼接
	query := fmt.Sprintf(" user like '%%%s' or id like '%%%s'", data["keywords"], data["keywords"])
	// 查询指定字段
	fields := []string{
		"id",
		"user",
		"photo_id",
		"sort",
		"state",
		"phone",
		"email",
		"name",
		"card",
		"introduction",
		"sex",
		"age",
		"region",
		"created_at",
	}
	// 排序
	order := "created_at desc"
	// 执行查询
	list, err := repositories.Admin().List(query, fields, order, data)
	if err != nil {
		panic(AdminConstant.GetAdminListFail)
	}
	return list
}
