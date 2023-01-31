package config

// MysqlConfig 数据库的配置
type MysqlConfig struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port uint `mapstructure:"port"`
	// 连接用户名
	User string `mapstructure:"user"`
	// 连接密码
	Password string `mapstructure:"password"`
	// 连接数据库名称
	Dbname string `mapstructure:"dbname"`
	// 最大连接数
	MaxOpenConn uint `mapstructure:"max_open_conn"`
	// 最大连接空闲数，建议和max_open_conn一致
	MaxIdleConn uint `mapstructure:"max_idle_conn"`
}
