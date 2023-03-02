package config

// AppConfig 项目全局的配置
type AppConfig struct {
	// 应用名称
	AppName string `mapstructure:"app_name"`
	// 路由前缀
	AppPrefix string `mapstructure:"app_prefix"`
	// 时区
	Timezone string `mapstructure:"timezone"`
	// 应用运行模式
	Mode string `mapstructure:"mode"`
	// 启动服务IP地址
	Host string `mapstructure:"host"`
	// 应用启动端口
	Port int `mapstructure:"port"`
	// 语言
	Language string `mapstructure:"language"`
	// mysql配置信息
	*MysqlConfig `mapstructure:"mysql"`
	// redis配置信息
	*RedisConfig `mapstructure:"redis"`
	// 日志配置信息
	*LogConfig `mapstructure:"log"`
	// jwt配置信息
	*JwtConfig `mapstructure:"jwt"`
	// 验证码配置
	*CaptchaConfig `mapstructure:"captcha"`
	// Kafka配置
	*KafkaConfig `mapstructure:"kafka"`
	// 短信配置
	*SmsConfig `mapstructure:"sms"`
}
