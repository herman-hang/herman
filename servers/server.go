package servers

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/kernel/core"
	"github.com/herman-hang/herman/kernel/log"
	middlewares2 "github.com/herman-hang/herman/middlewares"
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
	// 设置gin框架运行模式
	gin.SetMode(settings.Config.Mode)
	// 启动gin框架
	engine := gin.New()
	// 注册中间件
	engine.Use(log.GinLogger()).Use(middlewares2.CatchError()).Use(middlewares2.ServerHandler())
	// 初始化路由
	core.Engine = routers.InitRouter(engine)
	// 启动服务
	Run(host, port)
}

// ZapLogs 初始化日志
// @return void
func ZapLogs() {
	if err := log.InitZapLogs(settings.Config.Log, settings.Config.Mode); err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("Init zapLog failed:%v", err)))
	}
	core.Log = zap.S()
}

// Run 定义Server服务启动的方法
// @param string host 服务IP地址
// @param uint port 端口
// @return void
func Run(host string, port uint) {
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf(`
  _    _                                 
 | |  | |      Version: %v                          
 | |__| | ___ _ __ _ __ ___   __ _ _ __  
 |  __  |/ _ \ '__| '_ ' _ \ / _' | '_ \ 
 | |  | |  __/ |  | | | | | | (_| | | | |
 |_|  |_|\___|_|  |_| |_| |_|\__,_|_| |_|
                                         
 Server start on address: %v

`, color.GreenString(settings.Version), color.GreenString(serverAddr))

	server := &http.Server{
		Addr:    serverAddr,
		Handler: core.Engine,
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
	core.Log.Infof("Receive signals: %v", ch)
	_ = server.Shutdown(ctx)
}
