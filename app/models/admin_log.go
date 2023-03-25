package models

import (
	"gorm.io/gorm"
	"time"
)

// AdminLog 管理员日志结构体
type AdminLog struct {
	Id        uint           `json:"id" gorm:"column:id;primary_key;comment:管理员日志ID"`
	Type      uint8          `json:"type" gorm:"column:type;comment:日志类型（1登录日志，2操作日志）"`
	AdminId   uint           `json:"adminId" gorm:"column:admin_id;comment:管理员ID"`
	Ip        string         `json:"ip" gorm:"column:ip;comment:IP地址"`
	Path      string         `json:"path" gorm:"column:path;comment:请求路由"`
	Method    string         `json:"method" gorm:"column:method;comment:请求方法"`
	Remark    string         `json:"remark" gorm:"column:remark;comment:备注"`
	Code      uint16         `json:"code" gorm:"column:code;comment:响应状态码"`
	State     uint16         `json:"state" gorm:"column:state;comment:状态（1失败，2成功）"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (AdminLog) TableName() string {
	return "admin_log"
}
