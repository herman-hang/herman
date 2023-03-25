package models

import (
	"gorm.io/gorm"
	"time"
)

// Dictionary 数据字典表结构体
type Dictionary struct {
	Id        uint           `json:"id" gorm:"column:id;primary_key;comment:数据字典ID"`
	Name      string         `json:"name" gorm:"column:name;comment:数据字典名称"`
	Code      string         `json:"code" gorm:"column:code;comment:数据字典KEY"`
	Remark    string         `json:"remark" gorm:"column:remark;comment:备注"`
	State     uint8          `json:"state" gorm:"column:state;default:2;comment:状态(1禁用，2表示启用)"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置表名
func (Dictionary) TableName() string {
	return "dictionary"
}
