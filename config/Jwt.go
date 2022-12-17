package config

import "time"

// JwtConfig JWT配置
type JwtConfig struct {
	// 密钥
	JwtSecret string `mapstructure:"jwt_secret"`
	// token有效时间（单位：小时）
	EffectTime time.Duration `mapstructure:"effect_time"`
}
