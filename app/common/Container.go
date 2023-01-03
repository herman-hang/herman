package common

import (
	captchaService "github.com/TestsLing/aj-captcha-go/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Engine  *gin.Engine
	Log     *zap.SugaredLogger
	Db      *gorm.DB
	Redis   *redis.Client
	Captcha *captchaService.CaptchaServiceFactory
)

// NewContainer 全局容器
// @param *gin.Engine engine 上下文
// @param *zap.SugaredLogger log 日志对象
// @param *gorm.DB db DB对象
// @param *redis.Client redis Redis对象
// @param *captchaService.CaptchaServiceFactory captcha 验证码对象
func NewContainer(
	engine *gin.Engine,
	log *zap.SugaredLogger,
	db *gorm.DB,
	redis *redis.Client,
	captcha *captchaService.CaptchaServiceFactory,
) {
	Engine, Log, Db, Redis, Captcha = engine, log, db, redis, captcha
}
