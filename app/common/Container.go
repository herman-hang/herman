package common

import (
	"github.com/casbin/casbin/v2"
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
	Casbin *casbin.CachedEnforcer
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
	casbin *casbin.CachedEnforcer,
) {
	Engine, Log, Db, Redis, Once, Casbin = engine, log, db, redis, once, casbin
}
