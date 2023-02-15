package main

import (
	"github.com/herman/app/common"
	"github.com/herman/servers"
	"github.com/herman/servers/settings"
	"go.uber.org/zap"
)

// main 项目入口函数
func main() {
	// 进行配置文件的初始化
	if err := settings.InitConfig(); err != nil {
		zap.S().Fatalf("Init Config falied: %v\n", err)
	}

	s, err := servers.NewServer(settings.Config)
	if err != nil {
		zap.S().Fatalf("New Server falied: %v\n", err)
	}

	common.NewContainer(s.Engine, s.Log)

	s.Run()
}
