package user

import (
	"context"
	"fmt"
	"github.com/fp/fp-gin-framework/app/common"
	UserConstant "github.com/fp/fp-gin-framework/app/constants/user"
	"github.com/fp/fp-gin-framework/app/repositories"
	"github.com/fp/fp-gin-framework/app/utils"
	"github.com/go-redis/redis/v8"
)

// Login 用户登录
// @param map data 前端请求数据
// @return interface{} 返回一个token值
func Login(data map[string]interface{}) interface{} {
	user := repositories.User.GetUserInfo(fmt.Sprintf("%s", data["user"]))
	// 设置上下文
	ctx := context.Background()
	// 设置Redis错误密码的key
	key := fmt.Sprintf("user_password_error:%d", user.Id)

	// 获取错误登录次数
	errorNumber, err := common.Redis.Get(ctx, key).Int()

	// 判断是否登录次数过多
	if err != redis.Nil && errorNumber > 3 {
		panic(UserConstant.ErrorLoginOverload)
	}

	// 密码验证
	if !utils.ComparePasswords(user.Password, fmt.Sprintf("%s", data["password"])) {
		panic(UserConstant.PasswordError)
	}

	// 返回token
	return utils.GenerateToken(&utils.Claims{Uid: user.Id, Guard: "user"})
}
