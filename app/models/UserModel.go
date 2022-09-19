package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	User         string `json:"user"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Region       string `json:"region"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Introduction string `json:"introduction"`
	Status       string `json:"status"`
}

func GetUserInfo(user string) (*User, error) {
	var users User
	err := Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &users, nil
}
