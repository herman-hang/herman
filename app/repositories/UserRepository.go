package repositories

import (
	UserConstant "github.com/herman-hang/herman/app/constants/user"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// UserRepository 用户表仓储层
type UserRepository struct {
	BaseRepository
}

// User 实例化用户表仓储层
// @param *gorm.DB tx 事务
// @return UserRepository 返回用户表仓储层
func User(tx ...*gorm.DB) *UserRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &UserRepository{BaseRepository{Model: new(models.Users), Db: tx[0]}}
	}
	return &UserRepository{BaseRepository{Model: new(models.Users), Db: core.Db}}
}

// GetUserInfo 获取用户信息
// @param interface{} attributes 用户id或者用户user
// @return user 返回当前user用户的信息
func (base UserRepository) GetUserInfo(attributes interface{}) (user models.Users) {
	err := base.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(UserConstant.GetUserInfoFail)
	}

	return user
}
