package models

import (
	"gorm.io/gorm"
	"time"
)

// Role 角色表结构体
type Role struct {
	Id           uint           `json:"id" gorm:"column:id;primary_key;comment:角色ID"`
	Name         string         `json:"name" gorm:"column:name;comment:角色名称"`
	Role         string         `json:"role" gorm:"column:role;comment:角色KEY"`
	State        uint8          `json:"state" gorm:"column:state;default:2;comment:状态(1已停用,2已启用)"`
	Sort         uint           `json:"sort" gorm:"column:sort;default:0;comment:排序"`
	Introduction string         `json:"introduction" gorm:"column:introduction;comment:简介"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置角色表名
func (Role) TableName() string {
	return "roles"
}
