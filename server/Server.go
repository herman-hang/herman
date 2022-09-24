package server

import (
	"context"
	"fmt"
	"fp-back-user/app/middlewares"
	"fp-back-user/logs"
	"fp-back-user/routers"
	"fp-back-user/settings"
	"fp-back-user/storage/mysql"
	r "fp-back-user/storage/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
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
	Config   *settings.AppConfig // 全局的配置信息
	Engine   *gin.Engine         // 对应的gin的服务引擎
	Log      *zap.SugaredLogger  // 对应服务的log
	Db       *gorm.DB            // 数据库连接db
	Redis    *redis.Client       // redis
	Validate *validator.Validate // 验证器
}

// NewServer 初始化服务
func NewServer(config *settings.AppConfig) (*Server, error) {
	// 初始化日志
	if err := logs.InitZapLogs(config.LogConfig, config.Mode); err != nil {
		zap.S().Fatalf("Init ZapLog failed:%v", err)
	}
	zap.S().Info("Init ZapLogger Success!")

	// 初始化数据库
	db, err := mysql.InitGormDatabase(config.MysqlConfig)
	if err != nil {
		zap.S().Fatalf("Init Mysql failed:%v", err)
	}
	zap.S().Info("Init Mysql Success!")

	// 初始化redis
	rdb, err := r.InitRedisConfig(config.RedisConfig)
	if err != nil {
		zap.S().Fatalf("Init Redis Failed:%v", err)
	}

	zap.S().Info("Init Redis Success!")

	// 验证器定义
	validate := validator.New()

	gin.SetMode(config.Mode)
	e := gin.New()
	// 注册中间件
	e.Use(logs.GinLogger())
	e.Use(middlewares.CatchError())

	return &Server{
		Config:   config,
		Engine:   e,
		Log:      zap.S(),
		Db:       db,
		Redis:    rdb,
		Validate: validate,
	}, nil
}

// Run 定义Server服务启动的方法
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

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(2000)*time.Second)
	defer cancel()
	ch := <-sig
	s.Log.Infof("Receive Signals: %v", ch)
	_ = server.Shutdown(ctx)
}

// Close 定义Server服务注销的方法
func (s *Server) Close() {
	_ = s.Redis.Close()
	db := s.Db.DB()
	if db != nil {
		_ = db.Close()
	}
}
