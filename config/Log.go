package config

// LogConfig log的配置
type LogConfig struct {
	// 日志级别模式
	Level string `mapstructure:"level"`
	// 日志文件名
	FileName string `mapstructure:"filename"`
	// 单个日志文件大小(MB)，日志大小到达max_size就开始backup
	MaxSize uint `mapstructure:"max_size"`
	// 旧日志保存的最大天数，默认保存所有旧日志文件
	MaxAge uint `mapstructure:"max_age"`
	// 旧日志保存的最大数量，默认保存所有旧日志文件
	MaxBackups uint `mapstructure:"max_backups"`
	// backup的日志是否使用本地时间戳，默认使用UTC时间
	LocalTime bool `mapstructure:"local_time"`
	// 对backup的日志是否进行压缩，默认不压缩
	Compress bool `mapstructure:"compress"`
}
