package common

import (
	"fp-back-user/settings"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Config *settings.AppConfig
	Engine *gin.Engine
	Log    *zap.SugaredLogger
	Db     *gorm.DB
	Redis  *redis.Client
)

func NewContainer(
	config *settings.AppConfig,
	engine *gin.Engine,
	log *zap.SugaredLogger,
	db *gorm.DB,
	redis *redis.Client,
) {
	Config, Engine, Log, Db, Redis = config, engine, log, db, redis
}
