package models

import (
	"gorm.io/gorm"
	"time"
)

// Role 角色表结构体
type Role struct {
	Id           uint           `json:"id" gorm:"column:id;primary_key;type:uint(11);not null;comment:角色ID"`
	Name         string         `json:"name" gorm:"column:name;type:varchar(20);not null;comment:角色名称"`
	Role         string         `json:"role" gorm:"column:role;type:varchar(20);unique;not null;comment:角色KEY"`
	State        uint8          `json:"state" gorm:"column:state;type:tinyint(4);default:2;not null;comment:状态(1已停用,2已启用)"`
	Sort         uint           `json:"sort" gorm:"column:sort;type:uint(11);default:0;not null;comment:排序"`
	Introduction string         `json:"introduction" gorm:"column:introduction;type:text;comment:简介"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置角色表名
func (Role) TableName() string {
	return "roles"
}
