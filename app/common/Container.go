package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	Engine *gin.Engine
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Redis  *redis.Client
	Once   *sync.Once
)

// NewContainer 全局容器
// @param *gin.Engine engine 上下文
// @param *zap.SugaredLogger log 日志对象
// @param *gorm.DB db DB对象
// @param *redis.Client redis Redis对象
func NewContainer(
	engine *gin.Engine,
	log *zap.SugaredLogger,
	db *gorm.DB,
	redis *redis.Client,
	once *sync.Once,
) {
	Engine, Log, Db, Redis, Once = engine, log, db, redis, once
}
