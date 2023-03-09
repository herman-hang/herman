package config

// FileStorage 文件存储配置
type FileStorage struct {
	Drive string `mapstructure:"drive"`
	// 本地存储配置
	*Local `mapstructure:"local"`
	// 阿里云OSS存储配置
	*Oss `mapstructure:"oss"`
	// 腾讯云COS存储配置
	*Cos `mapstructure:"cos"`
	// 七牛云存储配置
	*Qiniu `mapstructure:"qiniu"`
}

// Local 本地存储配置
type Local struct {
	// 本地存储路径
	Path string `mapstructure:"path"`
}

// Oss 阿里云OSS存储配置
type Oss struct {
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

// Qiniu 七牛云存储配置
type Qiniu struct {
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

// Cos 腾讯云COS存储配置
type Cos struct {
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
