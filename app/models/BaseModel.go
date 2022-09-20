package models

import (
	s "fp-back-user/server"
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
}

// GetUserInfo 获取用户信息
func GetUserInfo(user string) (*User, error) {
	var users User
	err := s.Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &users, nil
}
