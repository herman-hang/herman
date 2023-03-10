package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app/common"
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
			"id":        fileInfo["id"],
			"file_name": fileInfo["fileName"],
			"file_type": fileInfo["fileType"],
			"file_ext":  fileInfo["fileExt"],
			"file_size": fileInfo["fileSize"],
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
	}, []string{"id", "file_name", "file_path", "hash"})
	if err != nil {
		panic(FileConstant.NotExist)
	}
	// 设置响应头
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info["file_name"]))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(info["file_path"].(string))
}

// Preview 图片预览
// @param ctx *gin.Context 上下文
// @param data map[string]interface{} 请求参数
// @return void
func Preview(ctx *gin.Context, data map[string]interface{}) {
	info, err := repositories.File().Find(map[string]interface{}{
		"id": data["id"],
	}, []string{"id", "file_name", "file_ext", "file_path", "hash"})
	common.Log.Debug(info)
	if err != nil {
		panic(FileConstant.NotExist)
	}
	ctx.File(info["file_path"].(string) + info["file_path"].(string))
}
