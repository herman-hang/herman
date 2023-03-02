package settings

import (
	"github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config = new(config.AppConfig)

// InitConfig 初始化配置
// @return void
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		zap.S().Fatalf("viper ReadInConfig failed, err: %v\n", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		zap.S().Fatalf("viper unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(ReloadConfig)
}

// ReloadConfig 加载配置回调
// @param in 加载配置回调
func ReloadConfig(in fsnotify.Event) {
	zap.S().Info("Configuration modified!")
	err := viper.Unmarshal(&Config)
	if err != nil {
		zap.S().Fatalf("viper unmarshal failed, err:%v\n", err)
		return
	} else {
		zap.S().Info("Configuration reloading succeeded!")
		return
	}
}
