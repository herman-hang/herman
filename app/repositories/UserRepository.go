package repositories

import (
	"github.com/fp/fp-gin-framework/app/common"
	UserConstant "github.com/fp/fp-gin-framework/app/constants/user"
	"github.com/fp/fp-gin-framework/app/models"
	"github.com/jinzhu/gorm"
)

// User 实例化结构体并重写BaseRepository
var User = &UserRepository{BaseRepository{Model: new(models.Users)}}

type UserRepository struct {
	BaseRepository
}

// GetUserInfo 获取用户信息
// @param string user 用户名
// @return *Users error 返回当前user用户的信息和错误信息
func (u UserRepository) GetUserInfo(user string) (*models.Users, error) {
	var users models.Users
	err := common.Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return &users, nil
}
