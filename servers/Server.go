package servers

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/bootstrap/log"
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
	common.Engine = e
	// 启动服务
	Run(host, port)
}

// ZapLogs 初始化日志
// @return void
func ZapLogs() {
	if err := log.InitZapLogs(settings.Config.LogConfig, settings.Config.Mode); err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("Init zapLog failed:%v", err)))
	}
	common.Log = zap.S()
}

// Run 定义Server服务启动的方法
// @param string host 服务IP地址
// @param uint port 端口
// @return void
func Run(host string, port uint) {
	// 初始化路由
	routers.InitRouter(common.Engine)
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf(`
  _    _                                 
 | |  | |                                
 | |__| | ___ _ __ _ __ ___   __ _ _ __  
 |  __  |/ _ \ '__| '_ ' _ \ / _' | '_ \ 
 | |  | |  __/ |  | | | | | | (_| | | | |
 |_|  |_|\___|_|  |_| |_| |_|\__,_|_| |_|
                                         
 Server start on address: %v

`, color.GreenString(serverAddr))
	server := &http.Server{
		Addr:    serverAddr,
		Handler: common.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			zap.S().Fatal(color.RedString(fmt.Sprintf("Failed to start server, %v", err)))
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2000)*time.Second)
	defer cancel()
	ch := <-sig
	common.Log.Infof("Receive signals: %v", ch)
	_ = server.Shutdown(ctx)
}
