package config

// Mysql 数据库的配置
type Mysql struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port int `mapstructure:"port"`
	// 连接用户名
	User string `mapstructure:"user"`
	// 连接密码
	Password string `mapstructure:"password"`
	// 连接数据库名称
	Dbname string `mapstructure:"dbname"`
	// 最大连接数
	MaxOpenConn int `mapstructure:"max_open_conn"`
	// 最大连接空闲数，建议和max_open_conn一致
	MaxIdleConn int `mapstructure:"max_idle_conn"`
}
