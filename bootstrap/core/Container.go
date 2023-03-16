package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Engine *gin.Engine
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Redis  *redis.Client
	Casbin *casbin.CachedEnforcer
)

// Debug 重新封装Debug方法
// @param args ...interface{}
func Debug(args ...interface{}) {
	Log.Debug(args)
}
