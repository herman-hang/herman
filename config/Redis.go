package config

// RedisConfig redis的配置
type RedisConfig struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port int `mapstructure:"port"`
	// 连接密码
	Password string `mapstructure:"password"`
	// 默认数据库，默认是0
	Db int `mapstructure:"db"`
	// 最大连接数
	PoolSize int `mapstructure:"pool_size"`
}
