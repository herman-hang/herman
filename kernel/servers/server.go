package servers

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/kernel/app"
	"github.com/herman-hang/herman/kernel/log"
	middleware "github.com/herman-hang/herman/middlewares"
	"github.com/herman-hang/herman/routers"
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
	gin.SetMode(app.Config.Mode)
	// 启动gin框架
	engine := gin.New()
	// 注册中间件
	engine.Use(log.GinLogger()).Use(middleware.CatchError())
	// 跨域配置判断
	if app.Config.Cores.IsOpen {
		engine.Use(middleware.Cors())
	}
	// 初始化路由
	app.Engine = routers.InitRouter(engine)
	// 启动服务
	Run(host, port)
}

// ZapLogs 初始化日志
// @return void
func ZapLogs() {
	if err := log.InitZapLogs(app.Config.Log, app.Config.Mode); err != nil {
		zap.S().Fatal(color.RedString(fmt.Sprintf("Init zapLog failed:%v", err)))
	}
	app.Log = zap.S()
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

`, color.GreenString(app.Version), color.GreenString(serverAddr))

	server := &http.Server{
		Addr:    serverAddr,
		Handler: app.Engine,
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
	app.Log.Infof("Receive signals: %v", ch)
	_ = server.Shutdown(ctx)
}
