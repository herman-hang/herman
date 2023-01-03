package servers

import (
	"context"
	"fmt"
	captchaService "github.com/TestsLing/aj-captcha-go/service"
	"github.com/fp/fp-gin-framework/app/middlewares"
	"github.com/fp/fp-gin-framework/bootstrap/mysql"
	r "github.com/fp/fp-gin-framework/bootstrap/redis"
	"github.com/fp/fp-gin-framework/config"
	"github.com/fp/fp-gin-framework/routers"
	"github.com/fp/fp-gin-framework/servers/captcha"
	"github.com/fp/fp-gin-framework/servers/log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server 定义服务所需要的组件
type Server struct {
	Config  *config.AppConfig                     // 全局的配置信息
	Engine  *gin.Engine                           // 对应的gin的服务引擎
	Log     *zap.SugaredLogger                    // 对应服务的log
	Db      *gorm.DB                              // 数据库连接db
	Redis   *redis.Client                         // redis
	Captcha *captchaService.CaptchaServiceFactory // 验证码
}

// NewServer 初始化服务
// @param *settings.AppConfig config 应用配置信息
// @return *Server error 返回服务结构体和错误信息
func NewServer(config *config.AppConfig) (*Server, error) {
	gin.SetMode(config.Mode)
	e := gin.New()
	// 注册中间件
	e.Use(log.GinLogger())
	e.Use(middlewares.CatchError())

	return &Server{
		Config:  config,
		Engine:  e,
		Log:     ZapLogs(config),
		Db:      GormDatabase(config),
		Redis:   Redis(config),
		Captcha: Captcha(config),
	}, nil
}

// ZapLogs 初始化日志
// @param *config.AppConfig config 应用配置信息
// @return *zap.SugaredLogger 返回日志实例
func ZapLogs(config *config.AppConfig) *zap.SugaredLogger {
	if err := log.InitZapLogs(config.LogConfig, config.Mode); err != nil {
		zap.S().Fatalf("Init ZapLog failed:%v", err)
	}
	zap.S().Info("Init ZapLogger Success!")
	return zap.S()
}

// GormDatabase 初始化数据库
// @param *config.AppConfig config 应用配置信息
// @return *gorm.DB db 返回数据库实例
func GormDatabase(config *config.AppConfig) (db *gorm.DB) {
	db, err := mysql.InitGormDatabase(config.MysqlConfig)
	if err != nil {
		zap.S().Fatalf("Init Mysql failed:%v", err)
	}
	zap.S().Info("Init Mysql Success!")
	return db
}

// Redis 初始化redis
// @param *config.AppConfig config 应用配置信息
// @return *redis.Client rdb 返回Redis实例
func Redis(config *config.AppConfig) (rdb *redis.Client) {
	rdb, err := r.InitRedisConfig(config.RedisConfig)
	if err != nil {
		zap.S().Fatalf("Init Redis Failed:%v", err)
	}

	zap.S().Info("Init Redis Success!")
	return rdb
}

// Captcha 初始化滑块验证码
// @param *config.AppConfig config 应用配置信息
// @return captchaFactory 返回验证码工厂实例
func Captcha(config *config.AppConfig) (captchaFactory *captchaService.CaptchaServiceFactory) {
	captchaFactory = captcha.InitCaptcha(config)
	zap.S().Info("Init Captcha Success!")
	return captchaFactory
}

// Run 定义Server服务启动的方法
// @param *Server s 服务结构体
func (s *Server) Run() {
	defer s.Close()
	// 初始化路由
	routers.InitRouter(s.Engine)
	serverAddr := fmt.Sprintf("%s:%d", "0.0.0.0", s.Config.Port)
	s.Log.Infof("Server Start on Address: %v", serverAddr)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: s.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			s.Log.Fatalf("Failed to start server, %v", err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2000)*time.Second)
	defer cancel()
	ch := <-sig
	s.Log.Infof("Receive Signals: %v", ch)
	_ = server.Shutdown(ctx)
}

// Close 定义Server服务注销的方法
// @param *Server s 服务结构体
func (s *Server) Close() {
	_ = s.Redis.Close()
	db := s.Db.DB()
	if db != nil {
		_ = db.Close()
	}
}
