package servers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/bootstrap/log"
	"github.com/herman-hang/herman/config"
	"github.com/herman-hang/herman/routers"
	"github.com/herman-hang/herman/servers/settings"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewServer 初始化服务
// @param string host 服务IP地址
// @param uint port 端口
// @return void
func NewServer(host string, port uint) {
	gin.SetMode(settings.Config.Mode)
	e := gin.New()
	// 注册中间件
	e.Use(log.GinLogger()).Use(middlewares.CatchError()).Use(middlewares.ServerHandler())
	// 初始化日志
	zapLog := ZapLogs(settings.Config)

	// 初始化容器
	common.NewContainer(e, zapLog)
	// 启动服务
	Run(host, port)
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
// @param string host 服务IP地址
// @param uint port 端口
// @return void
func Run(host string, port uint) {
	// 初始化路由
	routers.InitRouter(common.Engine)
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	common.Log.Infof("Server Start on Address: %v", serverAddr)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: common.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			common.Log.Fatalf("Failed to start server, %v", err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2000)*time.Second)
	defer cancel()
	ch := <-sig
	common.Log.Infof("Receive Signals: %v", ch)
	_ = server.Shutdown(ctx)
}
