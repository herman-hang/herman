package common

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Engine  *gin.Engine
	Log     *zap.SugaredLogger
	Db      *gorm.DB
	Redis   *redis.Client
	Casbin  *casbin.CachedEnforcer
	Context *gin.Context
)

// NewContainer 全局容器
// @param *gin.Engine engine Engine
// @param *zap.SugaredLogger log 日志对象
func NewContainer(
	engine *gin.Engine,
	log *zap.SugaredLogger,
) {
	Engine, Log = engine, log
}
