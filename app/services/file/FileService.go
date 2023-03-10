package file

import (
	"github.com/gin-gonic/gin"
	FileConstant "github.com/herman-hang/herman/app/constants/file"
	"github.com/herman-hang/herman/app/models"
	"github.com/herman-hang/herman/app/repositories"
	"mime/multipart"
)

// Storage 文件存储接口实现
type Storage interface {
	Upload(key string, content []byte) error
	Download(key string) ([]byte, error)
	Preview(key string) error
}

// Upload 文件上传
// @param ctx *gin.Context 上下文
// @param files []*multipart.FileHeader 文件对象切片
// @return existFileInfos []map[string]interface{} 文件切片信息
func Upload(ctx *gin.Context, files []*multipart.FileHeader) (existFileInfos []map[string]interface{}) {
	// 获取登录信息
	info, _ := ctx.Get("admin")
	admin := info.(*models.Admin)
	// 执行文件上传
	fileInfos, existFileInfos := Exec(files, admin.Id)
	for _, info := range fileInfos {
		// 保存文件信息
		fileInfo, err := repositories.File().Insert(info)
		if err != nil {
			panic(FileConstant.RecordFileFail)
		}
		existFileInfos = append(existFileInfos, map[string]interface{}{
			"id":       fileInfo["id"],
			"fileName": fileInfo["fileName"],
			"fileType": fileInfo["fileType"],
			"fileExt":  fileInfo["fileExt"],
			"fileSize": fileInfo["fileSize"],
		})
	}
	return existFileInfos
}

// Download 文件下载
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return void
func Download(ctx *gin.Context, data map[string]interface{}) {
	info, err := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "drive", "file_name", "file_path", "hash"})
	if err != nil {
		panic(FileConstant.NotExist)
	}
	// 返回文件流
	stream := adaptiveDownload(info)
	// 响应文件流
	response(ctx, stream, info["fileName"].(string))
}

// Preview 图片预览
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return void
func Preview(ctx *gin.Context, data map[string]interface{}) {
	info, err := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "drive", "file_name", "file_type", "file_ext", "file_path", "file_size", "hash"})
	if err != nil {
		panic(FileConstant.NotExist)
	}
	// 判断是否为图片
	if info["fileType"].(string) != "image/jpeg" &&
		info["fileType"].(string) != "image/png" &&
		info["fileType"].(string) != "image/gif" {
		panic(FileConstant.NotImage)
	}
	// 返回文件流
	stream := adaptiveDownload(info)
	// 响应文件流
	ctx.Header("Content-Type", info["fileType"].(string))
	ctx.Header("Connection", "keep-alive")
	response(ctx, stream, info["fileName"].(string))
}
