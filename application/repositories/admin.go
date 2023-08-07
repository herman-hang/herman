package repositories

import (
	AdminConstant "github.com/herman-hang/herman/application/constants/admin"
	"github.com/herman-hang/herman/application/models"
	"github.com/herman-hang/herman/kernel/core"
	"gorm.io/gorm"
)

// AdminRepository 管理员表仓储层
type AdminRepository struct {
	BaseRepository
}

// Admin 实例化管理员表仓储层
// @param *gorm.DB tx 事务
// @return AdminRepository 返回管理员表仓储层
func Admin(tx ...*gorm.DB) *AdminRepository {
	if len(tx) > 0 && tx[0] != nil {
		return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: tx[0]}}
	}

	return &AdminRepository{BaseRepository{Model: new(models.Admin), Db: core.Db()}}
}

// GetAdminInfo 获取管理员信息
// @param interface{} attributes 管理员id或者管理员user
// @return admin 返回当前管理员的信息
func (base AdminRepository) GetAdminInfo(attributes interface{}) (admin *models.Admin) {
	var err error
	switch attributes.(type) {
	case uint:
		err = base.Db.Where("id = ?", attributes).Find(&admin).Error
	case string:
		err = base.Db.Where("user = ?", attributes).Find(&admin).Error

	}
	if err != nil {
		panic(AdminConstant.GetAdminInfoFail)
	}

	return admin
}
