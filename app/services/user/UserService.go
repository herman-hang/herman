package user

import (
	"fmt"
	userConstant "github.com/fp/fp-gin-framework/app/constants/user"
	"github.com/fp/fp-gin-framework/app/repositories"
	"github.com/fp/fp-gin-framework/app/utils"
)

// Login 用户登录
// @param map data 前端请求数据
// @return interface{} 返回一个token值
func Login(data map[string]interface{}) interface{} {
	user := repositories.User.GetUserInfo(fmt.Sprintf("%s", data["user"]))

	// 密码验证
	if !utils.ComparePasswords(user.Password, fmt.Sprintf("%s", data["password"])) {
		panic(userConstant.PasswordError)
	}

	// 返回token
	return utils.GenerateToken(&utils.UserClaims{Uid: user.Id, Guard: "user"})
}
