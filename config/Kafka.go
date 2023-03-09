package config

// Kafka Kafka配置
type Kafka struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port int `mapstructure:"port"`
}
