package models

import (
	"github.com/fp/fp-gin-framework/app/common"
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	User         string     `json:"user"`
	Password     string     `json:"password"`
	Photo        string     `json:"photo"`
	Nickname     string     `json:"nickname"`
	Name         string     `json:"name"`
	Card         string     `json:"card"`
	Sex          string     `json:"sex"`
	Age          int        `json:"age"`
	Region       string     `json:"region"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	Introduction string     `json:"introduction"`
	Status       string     `json:"status"`
	SignOutIp    string     `json:"sign_out_ip"`
	SignTotal    string     `json:"sign_total"`
	SignOutAt    string     `json:"sign_out_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
}

// GetUserInfo 获取用户信息
// @param string user 用户名
// @return *Users error 返回当前user用户的信息和错误信息
func GetUserInfo(user string) (*Users, error) {
	var users Users
	err := common.Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err.Error())
	}

	return &users, nil
}
