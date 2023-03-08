package models

import (
	"gorm.io/gorm"
	"time"
)

// DictionaryDetail 数据字典表结构体
type DictionaryDetail struct {
	Id           uint           `json:"id" gorm:"column:id;primary_key;comment:明细值ID"`
	DictionaryId uint           `json:"dictionaryId" gorm:"column:dictionary_id;comment:数据字典ID"`
	Name         string         `json:"name" gorm:"column:name;comment:明细值名称"`
	Code         string         `json:"code" gorm:"column:code;comment:明细值KEY"`
	Value        string         `json:"value" gorm:"column:value;comment:明细值"`
	Remark       string         `json:"remark" gorm:"column:remark;comment:备注"`
	Sort         uint           `json:"sort" gorm:"column:sort;default:0;comment:排序"`
	State        uint8          `json:"state" gorm:"column:state;default:2;comment:状态(1禁用，2表示启用)"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:created_at;type:time;not null;comment:创建时间"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updated_at;type:time;not null;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:time;index;comment:删除时间"`
}

// TableName 设置表名
func (DictionaryDetail) TableName() string {
	return "dictionary_detail"
}
