package models

import (
	"gorm.io/gorm"
	"time"
)

// AdminRoleModel 管理员角色中间表模型结构体
type AdminRoleModel struct {
	Id        uint           `json:"id" gorm:"column:id;primary_key;comment:主键ID"`
	AdminId   uint           `json:"adminId" gorm:"column:admin_id;comment:管理员ID"`
	RoleKey   string         `json:"roleKey" gorm:"column:role_key;comment:角色KEY"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (AdminRoleModel) TableName() string {
	return "admin_role"
}
