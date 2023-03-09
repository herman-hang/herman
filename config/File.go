package config

// FileStorageConfig 文件存储配置
type FileStorageConfig struct {
	Drive string `mapstructure:"drive"`
	// 本地存储配置
	*LocalConfig `mapstructure:"local"`
	// 阿里云OSS存储配置
	*OssConfig `mapstructure:"oss"`
	// 腾讯云COS存储配置
	*CosConfig `mapstructure:"cos"`
}

// LocalConfig 本地存储配置
type LocalConfig struct {
	// 本地存储路径
	Path string `mapstructure:"path"`
}

// OssConfig 阿里云OSS存储配置
type OssConfig struct {
	// 阿里云OSS的AccessKeyID
	AccessKeyId string `mapstructure:"access_key_id"`
	// 阿里云OSS的AccessKeySecret
	AccessKeySecret string `mapstructure:"access_key_secret"`
	// 阿里云OSS的Bucket
	Bucket string `mapstructure:"bucket"`
	// 阿里云OSS的Endpoint
	Endpoint string `mapstructure:"endpoint"`
	// 阿里云OSS的外网域名
	Domain string `mapstructure:"domain"`
}

// QiniuConfig 七牛云存储配置
type QiniuConfig struct {
	// 七牛云的AccessKey
	AccessKey string `mapstructure:"access_key"`
	// 七牛云的SecretKey
	SecretKey string `mapstructure:"secret_key"`
	// 七牛云的Bucket
	Bucket string `mapstructure:"bucket"`
	// 七牛云的外网域名
	Domain string `mapstructure:"domain"`
	// 七牛云的Zone
	Zone string `mapstructure:"zone"`
	// 七牛云的UseHttps
	UseHttps string `mapstructure:"use_https"`
}

// CosConfig 腾讯云COS存储配置
type CosConfig struct {
	// 腾讯云COS的AppId
	AppId string `mapstructure:"app_id"`
	// 腾讯云COS的SecretId
	SecretId string `mapstructure:"secret_id"`
	// 腾讯云COS的SecretKey
	SecretKey string `mapstructure:"secret_key"`
	// 腾讯云COS的Bucket
	Bucket string `mapstructure:"bucket"`
	// 腾讯云COS的Region
	Region string `mapstructure:"region"`
	// 腾讯云COS的外网域名
	Domain string `mapstructure:"domain"`
	// 腾讯云COS的外网域名
	UseHttps string `mapstructure:"use_https"`
}
