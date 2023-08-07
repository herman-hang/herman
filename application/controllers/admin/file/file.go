package file

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/application"
	FileService "github.com/herman-hang/herman/application/services/admin/file"
	FileValidate "github.com/herman-hang/herman/application/validates/admin/file"
)

// UploadFile 上传文件
// @param ctx *gin.Context 上下文
// @return void
func UploadFile(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	files := FileValidate.Check(ctx)
	context.Json(FileService.Upload(ctx, files))
}

// DownloadFile 下载文件
// @param ctx *gin.Context 上下文
// @return void
func DownloadFile(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	FileService.Download(ctx, data)
}

// PreviewFile 图片预览
// @param ctx *gin.Context 上下文
// @return void
func PreviewFile(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	FileService.Preview(ctx, data)
}

// Prepare 获取分片上传方案
// @param ctx *gin.Context 上下文
// @return void
func Prepare(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	context.Json(FileService.Prepare(ctx, FileValidate.Prepare.Check(data)))
}

// ChunkUpload 分片上传
// @param ctx *gin.Context 上下文
// @return void
func ChunkUpload(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	toMap, file := FileValidate.ChunkCheck(ctx, data)
	FileService.ChunkUpload(toMap, file)
	context.Json(nil)
}

// MergeFile 合并文件
// @param ctx *gin.Context 上下文
// @return void
func MergeFile(ctx *gin.Context) {
	context := application.Request{Context: ctx}
	data := context.Params()
	FileService.Merge(FileValidate.Merge.Check(data))
	context.Json(nil)
}
