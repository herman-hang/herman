package config

// AppConfig 项目全局的配置
type AppConfig struct {
	// 应用名称
	AppName string `mapstructure:"app_name"`
	// 路由前缀
	AppPrefix string `mapstructure:"app_prefix"`
	// 应用版本（格式：主版本.子版本.修订版本）
	Version string `mapstructure:"version"`
	// 应用运行模式
	Mode string `mapstructure:"mode"`
	// 应用启动端口
	Port int `mapstructure:"port"`
	// mysql配置信息
	*MysqlConfig `mapstructure:"mysql"`
	// redis配置信息
	*RedisConfig `mapstructure:"redis"`
	// 日志配置信息
	*LogConfig `mapstructure:"log"`
	// jwt配置信息
	*JwtConfig `mapstructure:"jwt"`
}
