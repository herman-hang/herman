package admin

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/herman/app/common"
	AdminConstant "github.com/herman/app/constants/admin"
	"github.com/herman/app/repositories"
	"github.com/herman/app/utils"
	"time"
)

// Login 管理员登录
// @param map data 前端请求数据
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

	// 返回token
	return utils.GenerateToken(&utils.Claims{Uid: admin.Id, Guard: "admin"})
}
