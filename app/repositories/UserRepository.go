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
// @param interface{} attributes 用户id或者用户user
// @return userInfo 返回当前user用户的信息
func (u UserRepository) GetUserInfo(attributes interface{}) (userInfo *models.Users) {
	var users *models.Users
	err := common.Db.Where("id = @attributes OR user = @attributes", map[string]interface{}{"attributes": attributes}).First(&users).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return userInfo
}
