package models

import (
	"gorm.io/gorm"
	"time"
)

// Menu 菜单表结构体
type Menu struct {
	Id        uint           `json:"id" gorm:"column:id;type:uint(11);primary_key;not null;comment:菜单ID"`
	Pid       uint           `json:"pid" gorm:"column:pid;type:uint(11);not null;comment:菜单PID"`
	Name      string         `json:"name" gorm:"column:name;type:varchar(20);not null;comment:菜单名称"`
	Path      string         `json:"path" gorm:"column:name;type:varchar(100);not null;comment:路由PATH"`
	Method    string         `json:"method" gorm:"column:method;type:varchar(20);not null;comment:PATH的请求方法"`
	Sort      uint           `json:"sort" gorm:"column:sort;type:uint(11);default:0;not null;comment:排序"`
	State     uint8          `json:"state" gorm:"column:state;type:tinyint(4);default:2;not null;comment:状态(1已停用,2已启用)"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置菜单表名
func (Menu) TableName() string {
	return "menus"
}
