package main

import (
	"github.com/fp/fp-gin-framework/app/common"
	"github.com/fp/fp-gin-framework/servers"
	"github.com/fp/fp-gin-framework/servers/settings"
	"go.uber.org/zap"
	"sync"
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

	common.NewContainer(s.Engine, s.Log, s.Db, s.Redis, new(sync.Once))

	s.Run()
}
