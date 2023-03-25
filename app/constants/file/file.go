package file

const (
	MkdirFail            = "文件目录创建失败"
	MaxMemory            = "文件上传请求体超过限制"
	Empty                = "文件不能为空"
	CountZero            = 0
	MaxCount             = 10
	SurpassMaxCount      = "文件数量超过限制"
	CloseFileFail        = "文件关闭失败"
	ExtFail              = "文件扩展名非法"
	TypeFail             = "文件类型非法"
	NameFail             = "文件名非法"
	SizeFail             = "文件大小超过限制10MB"
	UploadFail           = "上传失败"
	Download             = "下载失败"
	CreateOSSClientFail  = "创建OSS客户机失败"
	GetOSSBucketFail     = "获取存储驱动失败"
	ConfigFileDriveError = "文件存储驱动配置错误"
	NewObjectFail        = "实例化存储对象失败"
	OpenFileFail         = "打开文件失败"
	RecordFileFail       = "记录文件失败"
	ReadFileFail         = "读取文件失败"
	NotExist             = "文件不存在"
	NotImage             = "请求文件不是图片，无法预览"
)
