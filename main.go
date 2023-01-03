package main

import (
	"fmt"
	"github.com/fp/fp-gin-framework/app/common"
	"github.com/fp/fp-gin-framework/servers"
	"github.com/fp/fp-gin-framework/servers/settings"
)

// main 项目入口函数
func main() {
	// 进行配置文件的初始化
	if err := settings.InitConfig(); err != nil {
		fmt.Printf("Init Config falied: %v\n", err)
		return
	}

	s, err := servers.NewServer(settings.Config)
	if err != nil {
		fmt.Println(err)
		return
	}

	common.NewContainer(s.Engine, s.Log, s.Db, s.Redis, s.Captcha)

	s.Run()
}
