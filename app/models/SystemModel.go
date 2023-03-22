package models

import (
	"gorm.io/gorm"
	"time"
)

// System 系统设置结构体
type System struct {
	Id          uint           `json:"id" gorm:"column:id;primary_key;comment:系统设置ID"`
	Name        string         `json:"name" gorm:"column:name;comment:网站名称"`
	Title       string         `json:"title" gorm:"column:title;comment:网站标题"`
	Description string         `json:"description" gorm:"column:description;comment:网站描述"`
	Keywords    string         `json:"keywords" gorm:"column:keywords;comment:网站关键词"`
	LogoFileId  uint           `json:"logo_file_id" gorm:"column:logo_file_id;comment:LOGO文件ID"`
	IcoFileId   uint           `json:"ico_file_id" gorm:"column:ico_file_id;comment:ICO文件ID"`
	Record      string         `json:"record" gorm:"column:record;comment:备案号"`
	Copyright   string         `json:"copyright" gorm:"column:copyright;comment:底部版权声明"`
	IsWebsite   uint8          `json:"is_website" gorm:"column:is_website;comment:网站开关"`
	Email       string         `json:"email" gorm:"column:email;comment:邮箱"`
	Telephone   string         `json:"telephone" gorm:"column:telephone;comment:电话"`
	Address     string         `json:"address" gorm:"column:address;comment:地址"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (System) TableName() string {
	return "system"
}
