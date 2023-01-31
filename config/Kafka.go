package config

// KafkaConfig Kafka配置
type KafkaConfig struct {
	// 连接IP地址
	Host string `mapstructure:"host"`
	// 连接端口号
	Port uint `mapstructure:"port"`
}
