package repositories

import (
	"github.com/fp/fp-gin-framework/app/common"
	UserConstant "github.com/fp/fp-gin-framework/app/constants/user"
	"github.com/fp/fp-gin-framework/app/models"
	"gorm.io/gorm"
)

// User 实例化结构体并重写BaseRepository
var User = &UserRepository{BaseRepository{Model: new(models.Users)}}

// UserRepository 用户表仓储层
type UserRepository struct {
	BaseRepository
}

// GetUserInfo 获取用户信息
// @param interface{} attributes 用户id或者用户user
// @return user 返回当前user用户的信息
func (u UserRepository) GetUserInfo(attributes interface{}) (user models.Users) {
	err := common.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return user
}
