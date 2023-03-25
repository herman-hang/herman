package config

// Sms Kafka配置
type Sms struct {
	// API接口
	Api string `mapstructure:"api"`
	// 用户名
	User string `mapstructure:"user"`
	// 密码
	Password string `mapstructure:"password"`
}
