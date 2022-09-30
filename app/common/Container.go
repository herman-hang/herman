package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Engine *gin.Engine
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Redis  *redis.Client
)

func NewContainer(
	engine *gin.Engine,
	log *zap.SugaredLogger,
	db *gorm.DB,
	redis *redis.Client,
) {
	Engine, Log, Db, Redis = engine, log, db, redis
}
