package models

import (
	"gorm.io/gorm"
	"time"
)

// Menu 菜单表结构体
type Menu struct {
	Id        uint           `json:"id" gorm:"column:id;primary_key;comment:菜单ID"`
	Pid       uint           `json:"pid" gorm:"column:pid;comment:菜单PID"`
	Name      string         `json:"name" gorm:"column:name;comment:菜单名称"`
	Path      string         `json:"path" gorm:"column:path;comment:路由PATH"`
	Method    string         `json:"method" gorm:"column:method;comment:PATH的请求方法"`
	Sort      uint           `json:"sort" gorm:"column:sort;default:0;comment:排序"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (Menu) TableName() string {
	return "menus"
}
