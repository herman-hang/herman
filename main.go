package main

import (
	"github.com/herman-hang/herman/app/common"
	"github.com/herman-hang/herman/servers"
	"github.com/herman-hang/herman/servers/settings"
	"go.uber.org/zap"
)

// main 项目入口函数
func main() {
	// 进行配置文件的初始化
	if err := settings.InitConfig(); err != nil {
		zap.S().Fatalf("Init Config falied: %v\n", err)
	}
	// 初始化服务
	s, err := servers.NewServer(settings.Config)
	if err != nil {
		zap.S().Fatalf("New Server falied: %v\n", err)
	}
	// 初始化容器
	common.NewContainer(s.Engine, s.Log)
	// 启动服务
	s.Run()
}
