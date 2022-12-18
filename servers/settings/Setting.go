package settings

import (
	"fmt"
	"github.com/fp/fp-gin-framework/config"
	"github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config = new(config.AppConfig)

// InitConfig 初始化配置
// @return error err 返回错误信息
func InitConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err = viper.ReadInConfig()
	if err != nil {
		zap.S().Errorf("viper.ReadInConfig failed, err: %v\n", err)
		return err
	}

	err = viper.Unmarshal(Config)
	if err != nil {
		zap.S().Errorf("viper unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(ReloadConfig)

	return
}

// ReloadConfig 加载配置回调
// @param fsnotify.Event in 加载配置回调
func ReloadConfig(in fsnotify.Event) {
	fmt.Println("配置文件修改了.....")

	err := viper.Unmarshal(Config)
	if err != nil {
		zap.S().Errorf("viper unmarshal failed, err:%v\n", err)
		return
	} else {
		fmt.Println("配置文件reload成功!")
		return
	}
}
