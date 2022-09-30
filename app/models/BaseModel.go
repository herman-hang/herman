package models

import (
	"fp-back-user/app/common"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
}

// GetUserInfo 获取用户信息
// @param string user 用户名
func GetUserInfo(user string) (*Users, error) {
	var users Users
	err := common.Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err.Error())
	}

	return &users, nil
}
