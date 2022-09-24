package common

import (
	"fp-back-user/settings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Config   *settings.AppConfig
	Engine   *gin.Engine
	Log      *zap.SugaredLogger
	Db       *gorm.DB
	Redis    *redis.Client
	Validate *validator.Validate
)

func NewContainer(
	config *settings.AppConfig,
	engine *gin.Engine,
	log *zap.SugaredLogger,
	db *gorm.DB,
	redis *redis.Client,
	validate *validator.Validate,
) {
	Config, Engine, Log, Db, Redis, Validate = config, engine, log, db, redis, validate
}
