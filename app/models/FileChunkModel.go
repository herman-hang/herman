package models

import (
	"gorm.io/gorm"
	"time"
)

// FileChunk 文件分片表结构体
type FileChunk struct {
	Id          uint           `json:"id" gorm:"column:id;primary_key;comment:分片文件ID"`
	FileId      uint           `json:"fileId" gorm:"column:file_id;comment:文件ID"`
	ChunkNumber uint           `json:"chunkNumber" gorm:"column:chunk_number;comment:文件ID"`
	ChunkSize   uint64         `json:"chunkSize" gorm:"column:chunk_size;comment:分片大小(单位byte)"`
	ChunkPath   string         `json:"chunkPath" gorm:"column:chunk_path;comment:分片路径"`
	Hash        string         `json:"hash" gorm:"column:hash;comment:分片Hash值"`
	State       uint8          `json:"state" gorm:"column:hash;comment:上传状态(1未上传，2已上传)"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index;comment:删除时间"`
}

// TableName 设置表名
func (FileChunk) TableName() string {
	return "file_chunks"
}
