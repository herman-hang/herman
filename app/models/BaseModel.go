package models

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
)

type Model struct {
	gorm.Model
}

// Make 将db,rdb映射到模型中便于数据操作
func Make(db *gorm.DB, rdb *redis.Client) {
	Db = db
	Redis = rdb
}
