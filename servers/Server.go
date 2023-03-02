package servers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/bootstrap/log"
	"github.com/herman-hang/herman/config"
	"github.com/herman-hang/herman/routers"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server 定义服务所需要的组件
type Server struct {
	Config *config.AppConfig  // 全局的配置信息
	Engine *gin.Engine        // 对应的gin的服务引擎
	Log    *zap.SugaredLogger // 对应服务的log
}

// NewServer 初始化服务
// @param *settings.AppConfig config 应用配置信息
// @return *Server error 返回服务结构体和错误信息
func NewServer(config *config.AppConfig) (*Server, error) {
	gin.SetMode(config.Mode)
	e := gin.New()
	// 注册中间件
	e.Use(log.GinLogger()).Use(middlewares.CatchError()).Use(middlewares.ServerHandler())

	zapLog := ZapLogs(config)

	return &Server{
		Config: config,
		Engine: e,
		Log:    zapLog,
	}, nil
}

// ZapLogs 初始化日志
// @param *config.AppConfig config 应用配置信息
// @return *zap.SugaredLogger 返回日志实例
func ZapLogs(config *config.AppConfig) *zap.SugaredLogger {
	if err := log.InitZapLogs(config.LogConfig, config.Mode); err != nil {
		zap.S().Fatalf("Init ZapLog failed:%v", err)
	}
	return zap.S()
}

// Run 定义Server服务启动的方法
// @param *Server s 服务结构体
func (s *Server) Run() {
	// 初始化路由
	routers.InitRouter(s.Engine)
	serverAddr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
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
