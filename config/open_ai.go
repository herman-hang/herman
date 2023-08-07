package config

type OpenAi struct {
	// SecretKey
	SecretKey string `mapstructure:"secret_key"`
	// 代理API
	ProxyApi string `mapstructure:"proxy_api"`
}
