package models

import (
	"fp-back-user/app/common"
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	User         string     `json:"user"`
	Password     string     `json:"password"`
	Nickname     string     `json:"nickname"`
	Sex          string     `json:"sex"`
	Age          int        `json:"age"`
	Region       string     `json:"region"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	Introduction string     `json:"introduction"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
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
