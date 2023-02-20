package models

import (
	"gorm.io/gorm"
	"time"
)

// AdminRoleModel 管理员角色中间表模型结构体
type AdminRoleModel struct {
	Id        uint           `json:"id" gorm:"column:id;type:uint(11);primary_key;comment:主键ID"`
	AdminId   uint           `json:"adminId" gorm:"column:admin_id;type:uint(11);index;comment:管理员ID"`
	RoleKey   string         `json:"roleKey" gorm:"column:role_key;type:varchar(20);index;comment:角色KEY"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置表名
func (AdminRoleModel) TableName() string {
	return "admin_role"
}
