package cobra

import (
	"github.com/herman-hang/herman/app/command"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 定义命令行
var rootCmd = &cobra.Command{Use: "herman"}

// 注册命令行
func init() {
	// 执行命令前初始化操作
	cobra.OnInitialize(settings.InitConfig, func() {
		if command.IsMigrate {
			// 数据库迁移
			_ = command.Migrate("up")
		}
	}, middlewares.Reload)
	// 启动服务命令注册
	rootCmd.AddCommand(command.StartServerCmd)
	// 注册数据库迁移
	rootCmd.AddCommand(command.MigrationCmd)
}

// Execute 执行命令
// @return void
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
