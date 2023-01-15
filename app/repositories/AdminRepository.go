package repositories

import (
	"github.com/fp/fp-gin-framework/app/common"
	AdminConstant "github.com/fp/fp-gin-framework/app/constants/admin"
	"github.com/fp/fp-gin-framework/app/models"
	"gorm.io/gorm"
)

// Admin 实例化结构体并重写BaseRepository
var Admin = &UserRepository{BaseRepository{Model: new(models.Admin)}}

type AdminRepository struct {
	BaseRepository
}

// GetAdminInfo 获取管理员信息
// @param interface{} attributes 管理员id或者管理员user
// @return userInfo 返回当前管理员的信息
func (u UserRepository) GetAdminInfo(attributes interface{}) (admin models.Admin) {
	err := common.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(AdminConstant.GetAdminInfoFail)
	}

	return admin
}
