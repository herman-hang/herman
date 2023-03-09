package models

import (
	"gorm.io/gorm"
	"time"
)

// File 文件表结构体
type File struct {
	Id        uint           `json:"id" gorm:"column:id;primary_key;comment:文件ID"`
	Drive     string         `json:"drive" gorm:"column:drive;comment:文件存储驱动(local:本地，oss:阿里云OSS，cos:腾讯云COS，qiniu:七牛云)"`
	CreatorId uint           `json:"creatorId" gorm:"column:creator_id;comment:创建者ID"`
	FileName  string         `json:"fileName" gorm:"column:file_name;comment:文件名"`
	FileType  string         `json:"fileType" gorm:"column:file_type;comment:文件类型"`
	FileExt   string         `json:"fileExt" gorm:"column:file_ext;comment:文件扩展名"`
	FilePath  string         `json:"filePath" gorm:"column:file_path;comment:文件路径"`
	Hash      string         `json:"hash" gorm:"column:file_path;comment:文件Hash值"`
	FileSize  int64          `json:"fileSize" gorm:"column:file_size;comment:文件大小"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (File) TableName() string {
	return "files"
}
