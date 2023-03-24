package cobra

import (
	"github.com/herman-hang/herman/app/command"
	"github.com/herman-hang/herman/app/middlewares"
	"github.com/herman-hang/herman/servers"
	"github.com/herman-hang/herman/servers/settings"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 定义命令行
var rootCmd = &cobra.Command{Use: "herman"}

// 注册命令行
func init() {
	// 执行命令前初始化操作
	cobra.OnInitialize(settings.InitConfig, servers.ZapLogs, func() {
		if command.IsMigrate {
			// 数据库迁移
			_ = command.Migrate("up")
		}
		// 如果执行的是数据库迁移命令，则不需要加载初始化操作
		if !command.MigrationStatus {
			middlewares.Reload()
		}
	})

	// 注册框架版本命令
	rootCmd.AddCommand(command.HermanVersionCmd)
	// 注册启动服务命令
	rootCmd.AddCommand(command.StartServerCmd)
	// 注册数据库迁移
	rootCmd.AddCommand(command.MigrationCmd)
	// 生成一个JWT令牌
	rootCmd.AddCommand(command.GenerateJwtCmd)
}

// Execute 执行命令
// @return void
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
