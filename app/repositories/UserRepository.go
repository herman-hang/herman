package repositories

import (
	"encoding/json"
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

// UserInfo 根据ID获取用户信息
// @param uint id 用户id
// @return userMap 返回当前用户ID的信息
func UserInfo(id uint) (userMap map[string]interface{}) {
	var users models.Users
	err := common.Db.Where("id = ?", id).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err.Error())
	}

	data, _ := json.Marshal(&users)
	userMap = make(map[string]interface{})
	_ = json.Unmarshal(data, &userMap)

	return userMap
}

// GetUserInfo 获取用户信息
// @param string user 用户名
// @return users err 返回当前user用户的信息和错误信息
func (u UserRepository) GetUserInfo(user string) (users *models.Users, err error) {
	err = common.Db.Where("user = ?", user).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return users, nil
}
