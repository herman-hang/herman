package config

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	Text           string `mapstructure:"text"`             // 水印文字
	CacheType      string `mapstructure:"cache_type"`       // 验证码使用的缓存类型驱动（目前只支持redis）
	CacheExpireSec int    `mapstructure:"cache_expire_sec"` // 缓存有效时间（单位：秒）
	ResourcePath   string `mapstructure:"resource_path"`    // 项目的绝对路径: 图片、字体等
}
