package main

import (
	"fmt"
	"fp-back-user/server"
	"fp-back-user/settings"
)

/*
* 项目的主要入口文件
 */
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
