package repositories

import (
	UserConstant "github.com/herman-hang/herman/app/constants/user"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/bootstrap/core"
	"gorm.io/gorm"
)

// UserRepository 用户表仓储层
type UserRepository struct {
	BaseRepository
}

// User 实例化用户表仓储层
// @return UserRepository 返回用户表仓储层
func User() *UserRepository {
	return &UserRepository{BaseRepository{Model: new(models.Users)}}
}

// GetUserInfo 获取用户信息
// @param interface{} attributes 用户id或者用户user
// @return user 返回当前user用户的信息
func (u UserRepository) GetUserInfo(attributes interface{}) (user models.Users) {
	err := core.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return user
}
