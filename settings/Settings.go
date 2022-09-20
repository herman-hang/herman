package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var Config = new(AppConfig)

// AppConfig 项目全局的配置
type AppConfig struct {
	AppName string `mapstructure:"app_name"`
	Version string `mapstructure:"version"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`

	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
	*JwtConfig   `mapstructure:"jwt"`
}

// MysqlConfig 数据库的配置
type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdsConn  int    `mapstructure:"max_ids_conn"`
}

// RedisConfig redis的配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// LogConfig log的配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// JwtConfig JWT配置
type JwtConfig struct {
	JwtSecret  string        `mapstructure:"jwt_secret"`
	EffectTime time.Duration `mapstructure:"effect_time"`
}

// InitConfig 初始化配置
func InitConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

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
