package repositories

import (
	"github.com/herman-hang/herman/app/common"
	AdminConstant "github.com/herman-hang/herman/app/constants/admin"
	"github.com/herman-hang/herman/app/models"
)

// AdminRepository 管理员表仓储层
type AdminRepository struct {
	BaseRepository
}

// Admin 实例化管理员表仓储层
// @return AdminRepository 返回管理员表仓储层
func Admin() *AdminRepository {
	return &AdminRepository{BaseRepository{Model: new(models.Admin)}}
}

// GetAdminInfo 获取管理员信息
// @param interface{} attributes 管理员id或者管理员user
// @return admin 返回当前管理员的信息
func (u AdminRepository) GetAdminInfo(attributes interface{}) (admin *models.Admin) {
	err := common.Db.Where("id = ?", attributes).Or("user = ?", attributes).Find(&admin).Error
	if err != nil {
		panic(AdminConstant.GetAdminInfoFail)
	}

	return admin
}
