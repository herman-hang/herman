package config

import "time"

// Jwt JWT配置
type Jwt struct {
	// 密钥
	JwtSecret string `mapstructure:"secret"`
	// token有效时间（单位：小时）
	EffectTime time.Duration `mapstructure:"effect_time"`
}
