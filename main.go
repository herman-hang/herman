package main

import (
	"fmt"
	"fp-back-user/server"
	"fp-back-user/settings"
)

// main 项目入口函数
func main() {
	// 进行配置文件的初始化
	if err := settings.InitConfig(); err != nil {
		fmt.Printf("Init Config falied: %v\n", err)
		return
	}
	s, err := server.NewServer(settings.Config)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Run()
}
