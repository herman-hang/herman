package config

// Cores 跨域配置
type Cores struct {
	// 是否支持跨域
	IsOpen bool `mapstructure:"is_open"`
	// 允许跨域的域名,多个域名用','逗号隔开或者使用通配符
	Origins string `mapstructure:"origins"`
}
