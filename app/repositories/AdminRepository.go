package repositories

import (
	"github.com/herman/app/common"
	AdminConstant "github.com/herman/app/constants/admin"
	"github.com/herman/app/models"
	"gorm.io/gorm"
)

// Admin 实例化结构体并重写BaseRepository
var Admin = &AdminRepository{BaseRepository{Model: new(models.Admin)}}

// AdminRepository 管理员表仓储层
type AdminRepository struct {
	BaseRepository
}

// GetAdminInfo 获取管理员信息
// @param interface{} attributes 管理员id或者管理员user
// @return admin 返回当前管理员的信息
func (u AdminRepository) GetAdminInfo(attributes interface{}) (admin *models.Admin) {
	err := common.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(AdminConstant.GetAdminInfoFail)
	}

	return admin
}
