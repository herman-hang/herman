package file

const (
	MkdirFail            = "文件目录创建失败"
	MaxMemory            = "文件上传请求体超过限制"
	Empty                = "文件不能为空"
	CountZero            = 0
	MaxCount             = 10
	SurpassMaxCount      = "文件数量超过限制"
	CloseFileFail        = "关闭文件失败"
	ExtFail              = "文件扩展名非法"
	TypeFail             = "文件类型非法"
	NameFail             = "文件名非法"
	SizeFail             = "文件大小超过限制10MB"
	UploadFail           = "上传失败"
	DownloadFail         = "下载失败"
	CreateOSSClientFail  = "创建OSS客户机失败"
	GetOSSBucketFail     = "获取存储驱动失败"
	ConfigFileDriveError = "文件存储驱动配置错误"
	NewObjectFail        = "实例化存储对象失败"
	OpenFileFail         = "打开文件失败"
	RecordFileFail       = "记录文件失败"
	ReadFileFail         = "读取文件失败"
	NotExist             = "文件不存在"
	PreviewFail          = "预览失败"
	NotImage             = "请求文件不是图片，无法预览"
	PrepareFail          = "获取分片上传方案失败"
	CreateFail           = "分片失败"
	NotUploadState       = 1
	UploadState          = 2
	ChunkExist           = "分片已经存在"
	ChunkNotExist        = "该文件没有找到分片"
	ChunkNotUpload       = "分片编号 %d 未上传"
	CreateFileFail       = "创建文件失败"
	MergeFail            = "合并文件失败"
	RemoveFileFail       = "删除文件失败"
	NotSupport           = "当前设置不支持分片上传"
	WriterFail           = "文件写入失败"
	EncodeJsonFail       = "JSON编码失败"
	UploadSuccess        = "上传成功"
)
