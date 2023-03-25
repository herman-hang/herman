package config

// App 项目全局的配置
type App struct {
	// 应用名称
	AppName string `mapstructure:"app_name"`
	// 路由前缀
	AppPrefix string `mapstructure:"app_prefix"`
	// 时区
	Timezone string `mapstructure:"timezone"`
	// 应用运行模式
	Mode string `mapstructure:"mode"`
	// 语言
	Language string `mapstructure:"language"`
	// mysql配置信息
	*Mysql `mapstructure:"mysql"`
	// redis配置信息
	*Redis `mapstructure:"redis"`
	// 日志配置信息
	*Log `mapstructure:"log"`
	// jwt配置信息
	*Jwt `mapstructure:"jwt"`
	// 验证码配置
	*Captcha `mapstructure:"captcha"`
	// Kafka配置
	*Kafka `mapstructure:"kafka"`
	// 短信配置
	*Sms `mapstructure:"sms"`
	// 文件存储配置
	*FileStorage `mapstructure:"storage"`
}
