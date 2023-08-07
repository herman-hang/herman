package user

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	UserConstant "github.com/herman-hang/herman/application/constants/admin/user"
	"github.com/herman-hang/herman/application/repositories"
	"github.com/herman-hang/herman/kernel/core"
	utils2 "github.com/herman-hang/herman/kernel/utils"
	"time"
)

// Login 用户登录
// @param map data 前端请求数据
// @return interface{} 返回一个token值
func Login(data map[string]interface{}) interface{} {
	user := repositories.User().GetUserInfo(fmt.Sprintf("%s", data["user"]))
	// 设置上下文
	ctx := context.Background()
	// 设置Redis错误密码的key
	key := fmt.Sprintf("user_password_error:%d", user.Id)

	// 获取错误登录次数
	errorNumber, err := core.Redis().Get(ctx, key).Int()

	// 判断是否登录次数过多
	if err != redis.Nil && errorNumber > UserConstant.LoginErrorLimitNumber {
		panic(UserConstant.ErrorLoginOverload)
	}

	// 密码验证
	if !utils2.ComparePasswords(user.Password, fmt.Sprintf("%s", data["password"])) {
		core.Redis().Set(ctx, key, errorNumber+UserConstant.Increment, time.Minute*UserConstant.KeyValidity)
		panic(UserConstant.PasswordError)
	}

	// 返回token
	return utils2.GenerateToken(&utils2.Claims{Uid: user.Id, Guard: "user"})
}
