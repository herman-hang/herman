package config

// RedisConfig redis的配置
type RedisConfig struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port uint `mapstructure:"port"`
	// 连接用户名
	UserName string `mapstructure:"user_name"`
	// 连接密码
	Password string `mapstructure:"password"`
	// 默认数据库，默认是0
	Db uint `mapstructure:"db"`
	// 最大连接数
	PoolSize uint `mapstructure:"pool_size"`
}
