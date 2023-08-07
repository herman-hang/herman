package cobra

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	command "github.com/herman-hang/herman/cmd"
	"github.com/herman-hang/herman/kernel/app"
	"github.com/herman-hang/herman/kernel/core"
	"github.com/herman-hang/herman/kernel/servers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

// rootCmd 定义命令行
var rootCmd = &cobra.Command{Use: "herman"}

// 注册命令行
func init() {
	// 执行命令前初始化操作
	cobra.OnInitialize(InitConfig, servers.ZapLogs, func() {
		if command.IsMigrate {
			// 数据库迁移
			_ = command.Migrate("up")
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

// InitConfig 初始化配置
// @return void
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(core.RootPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper ReadInConfig failed, err: %v", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&app.Config); err != nil {
		fmt.Printf("viper unmarshal failed, err:%v\n", err)
		os.Exit(1)
	}

	viper.WatchConfig()
	viper.OnConfigChange(ReloadConfig)
}

// ReloadConfig 加载配置回调
// @param in 加载配置回调
func ReloadConfig(in fsnotify.Event) {
	zap.S().Info("Configuration modified!")
	err := viper.Unmarshal(&app.Config)
	if err != nil {
		zap.S().Fatalf("viper unmarshal failed, err:%v\n", err)
		return
	} else {
		zap.S().Info("Configuration reloading succeeded!")
		return
	}
}
