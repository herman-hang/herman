package file

import (
	"github.com/gin-gonic/gin"
	"github.com/herman-hang/herman/app"
	FileService "github.com/herman-hang/herman/app/services/file"
	FileValidate "github.com/herman-hang/herman/app/validates/file"
)

// UploadFile 上传文件
// @param ctx *gin.Context 上下文
// @return void
func UploadFile(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	files := FileValidate.Check(ctx)
	context.Json(FileService.Upload(ctx, files))
}

// DownloadFile 下载文件
// @param ctx *gin.Context 上下文
// @return void
func DownloadFile(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	FileService.Download(ctx, data)
}

// PreviewFile 图片预览
// @param ctx *gin.Context 上下文
// @return void
func PreviewFile(ctx *gin.Context) {
	context := app.Request{Context: ctx}
	data := context.Params()
	FileService.Preview(ctx, data)
}
