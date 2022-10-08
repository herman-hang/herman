package user

import (
	"fmt"
	userConstant "fp-back-user/app/constants/user"
	"fp-back-user/app/models"
	"fp-back-user/app/utils"
)

// Login 用户登录
// @param map data 前端请求数据
func Login(data map[string]interface{}) interface{} {
	info, err := models.GetUserInfo(fmt.Sprintf("%v", data["user"]))
	if err != nil {
		panic(err.Error())
	}

	// 密码验证
	if !utils.ComparePasswords(info.Password, fmt.Sprintf("%v", data["password"])) {
		panic(userConstant.PasswordError)
	}

	// 返回token
	return utils.GenerateToken(&utils.UserClaims{ID: info.Id})
}
